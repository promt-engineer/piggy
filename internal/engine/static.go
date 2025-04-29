package engine

// Symbol represents a slot machine symbol
type Symbol int

// Constants for different symbols
const (
	None     Symbol = iota
	Dynamite        // 1
	Bat             // 2
	Saw             // 3
	Hammer          // 4
	Key             // 5
	A               // 6
	K               // 7
	Q               // 8
	J               // 9
	Bonus           // 10
	Wild            // Wild (Piggy)
)

// Multipliers for different symbols and combinations
var symbolMultipliers = map[Symbol]map[int]int64{
	Dynamite: {5: 200, 4: 60, 3: 30},
	Bat:      {5: 100, 4: 50, 3: 20},
	Saw:      {5: 60, 4: 25, 3: 10},
	Hammer:   {5: 50, 4: 20, 3: 10},
	Key:      {5: 25, 4: 15, 3: 5},
	A:        {5: 25, 4: 15, 3: 5},
	K:        {5: 15, 4: 10, 3: 5},
	Q:        {5: 15, 4: 10, 3: 5},
	J:        {5: 15, 4: 10, 3: 5},
}

// Position represents a position in the slot window
type Position struct {
	Col int
	Row int
}

// Define paylines for a 5x3 slot machine based on docs/paylines.csv
var Paylines = [][]Position{
	// Line 1: Middle row
	{
		{0, 1}, {1, 1}, {2, 1}, {3, 1}, {4, 1},
	},
	// Line 2: Top row
	{
		{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0},
	},
	// Line 3: Bottom row
	{
		{0, 2}, {1, 2}, {2, 2}, {3, 2}, {4, 2},
	},
	// Line 4: V shape
	{
		{0, 0}, {1, 1}, {2, 2}, {3, 1}, {4, 0},
	},
	// Line 5: Inverted V shape
	{
		{0, 2}, {1, 1}, {2, 0}, {3, 1}, {4, 2},
	},
	// Line 6: Zigzag
	{
		{0, 1}, {1, 0}, {2, 1}, {3, 0}, {4, 1},
	},
	// Line 7: Zigzag inverse
	{
		{0, 1}, {1, 2}, {2, 1}, {3, 2}, {4, 1},
	},
	// Line 8: Zigzag top-middle
	{
		{0, 0}, {1, 1}, {2, 0}, {3, 1}, {4, 0},
	},
	// Line 9: Zigzag bottom-middle
	{
		{0, 2}, {1, 1}, {2, 2}, {3, 1}, {4, 2},
	},
	// Line 10
	{
		{0, 1}, {1, 0}, {2, 0}, {3, 0}, {4, 1},
	},
	// Line 11
	{
		{0, 1}, {1, 2}, {2, 2}, {3, 2}, {4, 1},
	},
	// Line 12
	{
		{0, 2}, {1, 2}, {2, 1}, {3, 2}, {4, 2},
	},
	// Line 13
	{
		{0, 0}, {1, 0}, {2, 1}, {3, 0}, {4, 0},
	},
	// Line 14
	{
		{0, 2}, {1, 1}, {2, 1}, {3, 1}, {4, 2},
	},
	// Line 15
	{
		{0, 0}, {1, 1}, {2, 1}, {3, 1}, {4, 0},
	},
	// Line 16
	{
		{0, 0}, {1, 2}, {2, 0}, {3, 2}, {4, 0},
	},
	// Line 17
	{
		{0, 2}, {1, 0}, {2, 2}, {3, 0}, {4, 2},
	},
	// Line 18
	{
		{0, 1}, {1, 1}, {2, 0}, {3, 1}, {4, 1},
	},
	// Line 19
	{
		{0, 1}, {1, 1}, {2, 2}, {3, 1}, {4, 1},
	},
	// Line 20
	{
		{0, 2}, {1, 2}, {2, 0}, {3, 2}, {4, 2},
	},
	// Line 21
	{
		{0, 0}, {1, 0}, {2, 2}, {3, 0}, {4, 0},
	},
	// Line 22
	{
		{0, 0}, {1, 0}, {2, 1}, {3, 2}, {4, 2},
	},
	// Line 23
	{
		{0, 2}, {1, 2}, {2, 1}, {3, 0}, {4, 0},
	},
	// Line 24
	{
		{0, 1}, {1, 0}, {2, 2}, {3, 0}, {4, 1},
	},
	// Line 25
	{
		{0, 1}, {1, 2}, {2, 0}, {3, 2}, {4, 1},
	},
	// Line 26
	{
		{0, 1}, {1, 2}, {2, 1}, {3, 0}, {4, 0},
	},
	// Line 27
	{
		{0, 1}, {1, 0}, {2, 1}, {3, 2}, {4, 2},
	},
	// Line 28
	{
		{0, 0}, {1, 1}, {2, 2}, {3, 2}, {4, 2},
	},
	// Line 29
	{
		{0, 2}, {1, 1}, {2, 0}, {3, 0}, {4, 0},
	},
	// Line 30
	{
		{0, 0}, {1, 0}, {2, 0}, {3, 1}, {4, 2},
	},
	// Line 31
	{
		{0, 2}, {1, 2}, {2, 2}, {3, 1}, {4, 0},
	},
	// Line 32
	{
		{0, 1}, {1, 0}, {2, 1}, {3, 2}, {4, 1},
	},
	// Line 33
	{
		{0, 1}, {1, 2}, {2, 1}, {3, 0}, {4, 1},
	},
	// Line 34
	{
		{0, 0}, {1, 1}, {2, 1}, {3, 1}, {4, 1},
	},
	// Line 35
	{
		{0, 2}, {1, 1}, {2, 1}, {3, 1}, {4, 1},
	},
	// Line 36
	{
		{0, 0}, {1, 0}, {2, 1}, {3, 1}, {4, 1},
	},
	// Line 37
	{
		{0, 2}, {1, 2}, {2, 1}, {3, 1}, {4, 1},
	},
	// Line 38
	{
		{0, 2}, {1, 1}, {2, 2}, {3, 1}, {4, 0},
	},
	// Line 39
	{
		{0, 0}, {1, 1}, {2, 0}, {3, 1}, {4, 2},
	},
	// Line 40
	{
		{0, 1}, {1, 0}, {2, 0}, {3, 0}, {4, 0},
	},
	// Line 41
	{
		{0, 1}, {1, 2}, {2, 2}, {3, 2}, {4, 2},
	},
	// Line 42
	{
		{0, 0}, {1, 0}, {2, 0}, {3, 1}, {4, 0},
	},
	// Line 43
	{
		{0, 2}, {1, 2}, {2, 2}, {3, 1}, {4, 2},
	},
	// Line 44
	{
		{0, 0}, {1, 1}, {2, 0}, {3, 0}, {4, 0},
	},
	// Line 45
	{
		{0, 2}, {1, 1}, {2, 2}, {3, 2}, {4, 2},
	},
	// Line 46
	{
		{0, 1}, {1, 0}, {2, 1}, {3, 1}, {4, 1},
	},
	// Line 47
	{
		{0, 1}, {1, 2}, {2, 1}, {3, 1}, {4, 1},
	},
	// Line 48
	{
		{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 2},
	},
	// Line 49
	{
		{0, 2}, {1, 2}, {2, 2}, {3, 2}, {4, 0},
	},
	// Line 50
	{
		{0, 1}, {1, 1}, {2, 1}, {3, 0}, {4, 1},
	},
}

// Reels represents the set of reels in the slot machine
type Reels struct {
	Reels [][]Symbol
}

// Window represents the visible symbols in the slot machine
type Window struct {
	Symbols [][]Symbol
}

// Spin represents a single spin result
type Spin struct {
	Window       *Window
	Stops        []int
	Wager        int64
	Award        int64
	BaseAwardVal int64
}

// RNG interface for random number generation
type RNG interface {
	Rand(max uint64) (uint64, error)
}

// Gamble represents a gamble feature
type Gamble struct {
	// Empty for now as we're not implementing gamble features yet
}

// RestoringIndexes interface for restoring game state
type RestoringIndexes interface {
	IsShown(spin interface{}) bool
	Update(payload interface{}) error
}

type ReelsetData struct {
	Name             string
	Weight           int
	Probability      float64
	WildsProbability float64
	RTP              float64
}

var ReelsetWeights = []int{85, 7, 2, 6}

var AllReelsetData = []ReelsetData{
	{
		Name:             "Main Math1",
		Weight:           85,
		Probability:      0.85,
		WildsProbability: 0.0,
		RTP:              0.1155632751,
	},
	{
		Name:             "Main Math2",
		Weight:           7,
		Probability:      0.07,
		WildsProbability: 0.0,
		RTP:              0.07477371305,
	},
	{
		Name:             "Main Math3",
		Weight:           2,
		Probability:      0.02,
		WildsProbability: 1.0,
		RTP:              0.1690189141,
	},
	{
		Name:             "Main Math4",
		Weight:           6,
		Probability:      0.06,
		WildsProbability: 1.0,
		RTP:              0.09961106012,
	},
}

const TotalRTP = 0.1128199856
