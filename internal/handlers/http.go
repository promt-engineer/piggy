package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"piggy-bank/internal/app"
	"piggy-bank/internal/engine"
	"piggy-bank/internal/rng"
)

type Handler struct {
	spinFactory *engine.SpinFactory
	rngService  *rng.Service
}

func NewHandler(app *app.App) *Handler {
	rngService := app.GetRngService()

	reels := engine.RealisticReels()

	spinFactory := engine.NewSpinFactory(reels, rngService.GetClient())

	return &Handler{
		spinFactory: spinFactory,
		rngService:  rngService,
	}
}

type SpinResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
	Result  struct {
		Wager       int64             `json:"wager"`
		Award       int64             `json:"award"`
		Stops       []int             `json:"stops"`
		Symbols     [][]engine.Symbol `json:"symbols"`
		SymbolsText [][]string        `json:"symbols_text"`
	} `json:"result,omitempty"`
}

func symbolToString(symbol engine.Symbol) string {
	switch symbol {
	case engine.Dynamite:
		return "DYNAMITE"
	case engine.Bat:
		return "BAT"
	case engine.Saw:
		return "SAW"
	case engine.Hammer:
		return "HAMMER"
	case engine.Key:
		return "KEY"
	case engine.A:
		return "A"
	case engine.K:
		return "K"
	case engine.Q:
		return "Q"
	case engine.J:
		return "J"
	case engine.Bonus:
		return "BONUS"
	case engine.Wild:
		return "WILD"
	default:
		return "UNKNOWN"
	}
}

func (h *Handler) HandleSpin(w http.ResponseWriter, r *http.Request) {
	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	resp := SpinResponse{Success: true}

	wagerStr := r.URL.Query().Get("wager")
	if wagerStr == "" {
		resp.Success = false
		resp.Error = "missing wager parameter"
		json.NewEncoder(w).Encode(resp)
		return
	}

	wager, err := strconv.ParseInt(wagerStr, 10, 64)
	if err != nil {
		resp.Success = false
		resp.Error = "invalid wager value"
		json.NewEncoder(w).Encode(resp)
		return
	}

	spin, err := h.spinFactory.Generate(wager)
	if err != nil {
		resp.Success = false
		resp.Error = err.Error()
		json.NewEncoder(w).Encode(resp)
		return
	}

	resp.Result.Wager = spin.Wager
	resp.Result.Award = spin.Award
	resp.Result.Stops = spin.Stops

	resp.Result.Symbols = make([][]engine.Symbol, len(spin.Window.Symbols))
	resp.Result.SymbolsText = make([][]string, len(spin.Window.Symbols))

	for i, col := range spin.Window.Symbols {
		resp.Result.Symbols[i] = make([]engine.Symbol, len(col))
		resp.Result.SymbolsText[i] = make([]string, len(col))

		for j, symbol := range col {
			resp.Result.Symbols[i][j] = symbol
			resp.Result.SymbolsText[i][j] = symbolToString(symbol)
		}
	}

	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /spin", h.HandleSpin)
}

func SetupServer(address string, handler *Handler) *http.Server {
	mux := http.NewServeMux()
	handler.SetupRoutes(mux)

	server := &http.Server{
		Addr:    address,
		Handler: mux,
	}

	return server
}
