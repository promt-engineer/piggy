package app

import (
	"fmt"
	"log"
	"time"

	"piggy-bank/config"
	"piggy-bank/internal/rng"
)

type App struct {
	Config     *config.Config
	RngService *rng.Service
}

func NewApp(configPath string) (*App, error) {
	log.Printf("Initializing app with config path %s", configPath)
	startTime := time.Now()

	cfg, err := config.Load(configPath)
	if err != nil {
		return nil, fmt.Errorf("error loading config: %w", err)
	}
	log.Printf("Config loaded successfully %v", time.Since(startTime))

	log.Printf("Initializing RNG service...")
	rngService, err := rng.NewService(cfg)
	if err != nil {
		return nil, fmt.Errorf("error initializing RNG service: %w", err)
	}
	log.Printf("RNG service initialized successfully %v", time.Since(startTime))

	app := &App{
		Config:     cfg,
		RngService: rngService,
	}

	log.Printf("App initialized successfully %v", time.Since(startTime))
	return app, nil
}

func (a *App) GetConfig() *config.Config {
	return a.Config
}

func (a *App) GetRngService() *rng.Service {
	return a.RngService
}
