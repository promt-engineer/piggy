package engine

import (
	"fmt"
)

// NewReels creates a new set of reels
func NewReels() *Reels {
	return &Reels{
		Reels: [][]Symbol{
			{Dynamite, Bat, Saw, Hammer, Key, A, K, Q, J, Bonus, Wild},
			{Dynamite, Bat, Saw, Hammer, Key, A, K, Q, J, Bonus, Wild},
			{Dynamite, Bat, Saw, Hammer, Key, A, K, Q, J, Bonus, Wild},
			{Dynamite, Bat, Saw, Hammer, Key, A, K, Q, J, Bonus, Wild},
			{Dynamite, Bat, Saw, Hammer, Key, A, K, Q, J, Bonus, Wild},
		},
	}
}

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

	return &Spin{
		Window:       window,
		Stops:        stops,
		Wager:        wager,
		Award:        award,
		BaseAwardVal: award,
	}, nil
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
}

// Implement the Spin interface
func (s *Spin) BaseAward() int64 {
	return s.BaseAwardVal
}

func (s *Spin) BonusAward() int64 {
	return 0 // No bonus in this simple implementation
}

func (s *Spin) GetWager() int64 {
	return s.Wager
}

func (s *Spin) OriginalWager() int64 {
	return s.Wager
}

func (s *Spin) BonusTriggered() bool {
	return false // No bonus in this simple implementation
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
