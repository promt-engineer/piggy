package engine

import (
	"fmt"
)

// NewWindow creates a new window with the given dimensions
func NewWindow(width, height int) *Window {
	symbols := make([][]Symbol, width)
	for i := range symbols {
		symbols[i] = make([]Symbol, height)
	}
	return &Window{
		Symbols: symbols,
	}
}

// NewCoinStore creates a new coin store
func NewCoinStore() *CoinStore {
	return &CoinStore{
		CoinsCount:  0,
		FGTriggered: false,
		FGSpins:     0,
	}
}

// SpinFactory handles creating spins
type SpinFactory struct {
	reels    *Reels
	rng      RNG
	reelsets []*Reels
}

// NewSpinFactory creates a new spin factory
func NewSpinFactory(reels *Reels, rng RNG) *SpinFactory {
	// Create an array of reelsets
	reelsets := []*Reels{reel1, reel2, reel3, reel4}

	return &SpinFactory{
		reels:    reels,
		rng:      rng,
		reelsets: reelsets,
	}
}

// Generate creates a new spin
func (s *SpinFactory) Generate(wager int64) (*Spin, error) {
	if wager <= 0 {
		return nil, fmt.Errorf("wager must be positive")
	}

	// Select a reelset based on weights
	selectedReels, reelsetIndex, err := selectReelset(s.rng)
	if err != nil {
		return nil, fmt.Errorf("failed to select reelset: %w", err)
	}

	// Generate random stops
	stops := make([]int, len(selectedReels.Reels))
	for i := range stops {
		val, err := s.rng.Rand(uint64(len(selectedReels.Reels[i])))
		if err != nil {
			return nil, fmt.Errorf("failed to generate random number: %w", err)
		}
		stops[i] = int(val)
	}

	// Create window from stops
	window := NewWindow(len(selectedReels.Reels), 3)
	for i, stop := range stops {
		for j := 0; j < 3; j++ {
			symbolIndex := (stop + j) % len(selectedReels.Reels[i])
			window.Symbols[i][j] = selectedReels.Reels[i][symbolIndex]
		}
	}

	if reelsetIndex >= 0 && reelsetIndex < len(AllReelsetData) {
		reelsetData := AllReelsetData[reelsetIndex]
		if reelsetData.WildsProbability > 0 {
			// Для каждой позиции проверяем шанс замены на дикий символ
			for i := range window.Symbols {
				for j := range window.Symbols[i] {
					// Пропускаем, если уже дикий символ или бонусный символ
					if window.Symbols[i][j] == Wild || window.Symbols[i][j] == Bonus {
						continue
					}

					// Проверяем шанс замены на дикий символ
					wildChance, err := s.rng.Rand(100)
					if err == nil && float64(wildChance)/100.0 < reelsetData.WildsProbability {
						window.Symbols[i][j] = Wild
					}
				}
			}
		}
	}

	// Calculate award
	award := s.calculateAward(window, wager)

	// Create spin result
	spin := &Spin{
		Window:       window,
		Stops:        stops,
		Wager:        wager,
		Award:        award,
		BaseAwardVal: award,
	}

	return spin, nil
}

// GenerateWithCoins creates a new spin and handles coin accumulation
func (s *SpinFactory) GenerateWithCoins(wager int64, coins *CoinStore) (*Spin, error) {
	if wager <= 0 {
		return nil, fmt.Errorf("wager must be positive")
	}

	// Generate base spin
	spin, err := s.Generate(wager)
	if err != nil {
		return nil, err
	}

	// Add coin store reference
	spin.Coins = coins

	// Check for wild symbols in the window and add coins
	spin.CoinsAdded = s.processCoinAccumulation(spin)

	// Check if free game is triggered
	if spin.CoinsAdded > 0 {
		spin.FGTriggered = s.checkFreeGameTrigger(spin)
	}

	return spin, nil
}

// processCoinAccumulation checks for wild symbols and adds coins
func (s *SpinFactory) processCoinAccumulation(spin *Spin) int {
	// Skip if this is already a free game
	if spin.IsFreeGame {
		return 0
	}

	// Check for Wild symbols in the window
	hasWild := false
	for i := range spin.Window.Symbols {
		for j := range spin.Window.Symbols[i] {
			if spin.Window.Symbols[i][j] == Wild {
				hasWild = true
				break
			}
		}
		if hasWild {
			break
		}
	}

	// If no wild symbols, return 0
	if !hasWild {
		return 0
	}

	// Determine number of coins to add (1 or 2) based on weights
	val, err := s.rng.Rand(100)
	if err != nil {
		return 0
	}

	// Decide based on CoinCountWeights (80% for 1 coin, 20% for 2 coins)
	cumulativeWeight := uint64(0)
	coinsToAdd := 0
	for i, weight := range CoinCountWeights {
		cumulativeWeight += uint64(weight)
		if val < cumulativeWeight {
			coinsToAdd = i + 1 // Add 1 or 2 coins
			break
		}
	}

	// Update coin count
	if spin.Coins != nil {
		spin.Coins.CoinsCount += coinsToAdd
	}

	return coinsToAdd
}

// checkFreeGameTrigger checks if free game is triggered based on current coin count
func (s *SpinFactory) checkFreeGameTrigger(spin *Spin) bool {
	if spin.Coins == nil || spin.Coins.FGTriggered {
		return false
	}

	// Helper function to check trigger probability for a single increment
	checkTrigger := func() bool {
		coinCount := spin.Coins.CoinsCount

		// Check if we have enough coins to potentially trigger and if we haven't triggered already
		if coinCount < 8 {
			return false
		}

		// Get trigger probability based on coin count (capped at 17 coins)
		triggerProb := 0
		if coinCount > 17 {
			triggerProb = 100 // 100% trigger at >17 coins
		} else {
			triggerProb = FGTriggerProbabilities[coinCount]
		}

		// Roll for triggering
		val, err := s.rng.Rand(100)
		if err != nil {
			return false
		}

		return int(val) < triggerProb
	}

	// Check triggers based on number of coins added
	leftTrigger := false
	rightTrigger := false

	// First coin check
	if spin.CoinsAdded > 0 {
		leftTrigger = checkTrigger()
	}

	// Second coin check (if applicable)
	if spin.CoinsAdded > 1 {
		// Temporarily decrement coin count to check the second coin separately
		spin.Coins.CoinsCount--
		rightTrigger = checkTrigger()
		spin.Coins.CoinsCount++ // Restore the count
	}

	// Set number of free spins based on trigger results
	if leftTrigger && rightTrigger {
		spin.Coins.FGSpins = 12
		spin.Coins.FGTriggered = true
		return true
	} else if leftTrigger || rightTrigger {
		spin.Coins.FGSpins = 6
		spin.Coins.FGTriggered = true
		return true
	}

	return false
}

// GenerateFreeGame generates a spin for the free game mode
func (s *SpinFactory) GenerateFreeGame(wager int64, coins *CoinStore) (*Spin, error) {
	if wager <= 0 {
		return nil, fmt.Errorf("wager must be positive")
	}

	// Generate a regular spin
	spin, err := s.Generate(wager)
	if err != nil {
		return nil, err
	}

	// Mark as free game
	spin.IsFreeGame = true
	spin.Coins = coins

	// Free Game spins might have different multipliers or other special rules
	// that can be applied here

	return spin, nil
}

// calculateAward calculates the award for a window
func (s *SpinFactory) calculateAward(window *Window, wager int64) int64 {
	return s.calculateAwardWithPaylines(window, wager)
}

// calculateAwardWithPaylines calculates the award based on paylines
func (s *SpinFactory) calculateAwardWithPaylines(window *Window, wager int64) int64 {
	totalAward := int64(0)

	// Check each payline
	for _, payline := range Paylines {
		// Get symbols on this payline
		symbols := make([]Symbol, len(payline))
		for i, pos := range payline {
			if pos.Col < len(window.Symbols) && pos.Row < len(window.Symbols[pos.Col]) {
				symbols[i] = window.Symbols[pos.Col][pos.Row]
			}
		}

		// Count consecutive symbols from left to right
		award := s.evaluateSymbolLine(symbols, wager)
		totalAward += award
	}

	return totalAward
}

// evaluateSymbolLine evaluates a line of symbols for wins
func (s *SpinFactory) evaluateSymbolLine(symbols []Symbol, wager int64) int64 {
	if len(symbols) == 0 {
		return 0
	}

	// Find the first non-wild symbol (if any)
	targetSymbol := None
	for _, sym := range symbols {
		if sym != Wild && sym != None {
			targetSymbol = sym
			break
		}
	}

	// If no non-wild symbols found, use wild
	if targetSymbol == None {
		targetSymbol = Wild
	}

	// Count consecutive matching symbols from left
	count := 0
	for i := 0; i < len(symbols); i++ {
		if symbols[i] == targetSymbol || symbols[i] == Wild {
			count++
		} else {
			break
		}
	}

	// If we have at least 3 matching symbols, calculate win
	if count >= 3 {
		if multipliers, ok := symbolMultipliers[targetSymbol]; ok {
			if multiplier, ok := multipliers[count]; ok {
				return multiplier * wager / 100
			}
		}
	}

	return 0
}

// Weights for selecting reelsets - moved to static.go as ReelsetWeights
// var reelsetWeights = []int{85, 7, 2, 6}

// selectReelset selects a reelset based on weights from pick-probabilities.csv
// Returns the selected reelset and its index
func selectReelset(rng RNG) (*Reels, int, error) {
	// ВРЕМЕННАЯ МОДИФИКАЦИЯ: всегда возвращаем только reel1
	return reel4, 0, nil

	// Оригинальный код (закомментирован на время тестирования)
	/*
		// Get a random number from 0 to 99
		val, err := rng.Rand(100)
		if err != nil {
			return nil, -1, err
		}

		// Select based on weights
		cumulativeWeight := uint64(0)
		for i, weight := range ReelsetWeights {
			cumulativeWeight += uint64(weight)
			if val < cumulativeWeight {
				switch i {
				case 0:
					return reel1, 0, nil
				case 1:
					return reel2, 1, nil
				case 2:
					return reel3, 2, nil
				case 3:
					return reel4, 3, nil
				}
			}
		}

		// Default to the first reelset
		return reel1, 0, nil
	*/
}

// Implement the Spin interface
func (s *Spin) BaseAward() int64 {
	return s.BaseAwardVal
}

func (s *Spin) BonusAward() int64 {
	if s.IsFreeGame {
		return s.Award // In free game, all award is considered bonus
	}
	return 0
}

func (s *Spin) GetWager() int64 {
	return s.Wager
}

func (s *Spin) OriginalWager() int64 {
	return s.Wager
}

func (s *Spin) BonusTriggered() bool {
	return s.FGTriggered
}

func (s *Spin) DeepCopy() interface{} {
	newSpin := &Spin{
		Window: &Window{
			Symbols: make([][]Symbol, len(s.Window.Symbols)),
		},
		Stops:        make([]int, len(s.Stops)),
		Wager:        s.Wager,
		Award:        s.Award,
		BaseAwardVal: s.BaseAwardVal,
		IsFreeGame:   s.IsFreeGame,
		CoinsAdded:   s.CoinsAdded,
		FGTriggered:  s.FGTriggered,
	}

	// Deep copy the coin store if it exists
	if s.Coins != nil {
		newSpin.Coins = &CoinStore{
			CoinsCount:  s.Coins.CoinsCount,
			FGTriggered: s.Coins.FGTriggered,
			FGSpins:     s.Coins.FGSpins,
		}
	}

	for i, stop := range s.Stops {
		newSpin.Stops[i] = stop
	}

	for i, col := range s.Window.Symbols {
		newSpin.Window.Symbols[i] = make([]Symbol, len(col))
		copy(newSpin.Window.Symbols[i], col)
	}

	return newSpin
}

func (s *Spin) GetGamble() *Gamble {
	return nil // No gamble in this simple implementation
}

func (s *Spin) CanGamble(ri RestoringIndexes) bool {
	return false // No gamble in this simple implementation
}

// SimpleRestoringIndexes is a basic implementation of RestoringIndexes
type SimpleRestoringIndexes struct {
	IsShownVal bool
}

func (r *SimpleRestoringIndexes) IsShown(spin interface{}) bool {
	return r.IsShownVal
}

func (r *SimpleRestoringIndexes) Update(payload interface{}) error {
	// For simplicity, we'll just set IsShownVal to true
	r.IsShownVal = true
	return nil
}

// NewSimpleRestoringIndexes creates a new SimpleRestoringIndexes
func NewSimpleRestoringIndexes() *SimpleRestoringIndexes {
	return &SimpleRestoringIndexes{
		IsShownVal: false,
	}
}

// RealisticReels creates a more realistic set of reels for the slot machine
func RealisticReels() *Reels {
	return reel1
}

func GetAllReelsets() []*Reels {
	return []*Reels{reel1, reel2, reel3, reel4}
}

func NewSpinFactoryWithAllReelsets(rng RNG) *SpinFactory {
	return &SpinFactory{
		reels:    reel1,
		rng:      rng,
		reelsets: []*Reels{reel1, reel2, reel3, reel4},
	}
}

// FreeGameInfo returns information about the free game if triggered
func (s *Spin) FreeGameInfo() (triggered bool, spins int) {
	if s.Coins == nil {
		return false, 0
	}
	return s.FGTriggered, s.Coins.FGSpins
}

// IsFG returns whether this is a free game spin
func (s *Spin) IsFG() bool {
	return s.IsFreeGame
}

// GetCoins returns the current coin count
func (s *Spin) GetCoins() int {
	if s.Coins == nil {
		return 0
	}
	return s.Coins.CoinsCount
}
