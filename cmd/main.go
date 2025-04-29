package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"piggy-bank/internal/app"
	"piggy-bank/internal/handlers"
	"piggy-bank/internal/simulator"
)

func main() {
	application, err := app.NewApp("config.yaml")
	if err != nil {
		log.Fatalf("Error initializing app: %v", err)
	}

	cfg := application.GetConfig()

	addr := flag.String("addr", fmt.Sprintf(":%d", cfg.Server.Port), "HTTP server address")
	sim := flag.Bool("simulate", false, "Run simulation mode")

	flag.Parse()

	if *sim {
		runSimulation(application, cfg.Simulator.Spins, cfg.Simulator.Wager, cfg.Simulator.Workers, cfg.Simulator.ReportPath)
	} else {
		startServer(application, *addr)
	}
}

func startServer(app *app.App, address string) {
	log.Printf("Setting up HTTP server on %s", address)

	handler := handlers.NewHandler(app)
	server := handlers.SetupServer(address, handler)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Printf("Starting server on %s", address)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	<-done
	log.Print("Server shutdown initiated...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	log.Print("Server shutdown completed")
}

func runSimulation(app *app.App, spins, wager int64, workers int, outputPath string) {
	fmt.Printf("Starting simulation with %d spins, wager %d, using %d workers\n", spins, wager, workers)

	rngService := app.GetRngService()

	result, err := simulator.Simulate("piggy-bank", spins, wager, workers, rngService)
	if err != nil {
		log.Fatalf("Simulation failed: %v", err)
	}

	if err := os.MkdirAll(outputPath, 0o755); err != nil {
		log.Fatalf("Failed to create output directory: %v", err)
	}

	timestamp := time.Now().Format("2006-01-02-15-04-05")
	filename := fmt.Sprintf("piggy-bank-sim-%s.json", timestamp)
	fullPath := filepath.Join(outputPath, filename)

	jsonData, err := json.MarshalIndent(result.View(), "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal simulation results: %v", err)
	}

	// Write to file
	if err := os.WriteFile(fullPath, jsonData, 0o644); err != nil {
		log.Fatalf("Failed to write simulation results: %v", err)
	}

	// Also print summary to console
	view := result.View()
	fmt.Println("\n=== Simulation Results ===")
	fmt.Printf("Game: %s\n", view.Game)
	fmt.Printf("Spins: %s\n", view.Count)
	fmt.Printf("Wager: %s\n", view.Wager)
	fmt.Printf("Total Spent: %s\n", view.Spent)
	fmt.Printf("Max Exposure: %s\n", view.MaxExposure)
	fmt.Printf("RTP: %s%%\n", view.RTP)
	fmt.Printf("Hit Rate: %s\n", view.AwardRate)
	fmt.Printf("Volatility: %.3f\n", view.Volatility)
	fmt.Printf("\nDetailed report saved to: %s\n", fullPath)
}
