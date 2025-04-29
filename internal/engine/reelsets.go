package engine

// Reel sets for Piggy Bank game
// These reelsets are based on the CSV files provided

// reel1 - First reelset configuration with 81 symbols per reel
var reel1 = &Reels{
	Reels: [][]Symbol{
		// Reel 1
		{Dynamite, Dynamite, Dynamite, J, Key, K, Hammer, Q, Bat, Bat, Bat, K, K, K, Saw, J, Hammer, Q, Key, Key, Key, A, A, A, Hammer, J, J, J, Saw, Q, Bat, K, Dynamite, A, Saw, Saw, Saw, Q, Q, Q, Hammer, K, Key, Q, Saw, J, Key, K, Hammer, Q, Dynamite, A, A, A, Bat, J, Saw, Q, Hammer, K, Key, A, Hammer, J, Key, K, Bat, J, Dynamite, Q, Hammer, J, Hammer, A, A, Key, Q, Saw, K, Bat, J},
		// Reel 2
		{J, Bat, K, Saw, Q, Key, A, A, Hammer, J, Hammer, Q, Dynamite, J, Bat, K, Key, J, Hammer, A, Key, K, Hammer, Q, Saw, J, Bat, A, A, A, Dynamite, Q, Hammer, K, Key, J, Saw, Q, Key, K, Hammer, Q, Q, Q, Saw, Saw, Saw, A, Dynamite, K, Bat, Q, Saw, J, J, J, Hammer, A, A, A, Key, Key, Key, Q, Hammer, J, Saw, K, K, K, Bat, Bat, Bat, Q, Hammer, K, Key, J, Dynamite, Dynamite, Dynamite},
		// Reel 3
		{Dynamite, Dynamite, Dynamite, J, Key, K, Hammer, Q, Bat, Bat, Bat, K, K, K, Saw, J, Hammer, Q, Key, Key, Key, A, A, A, Hammer, J, J, J, Saw, Q, Bat, K, Dynamite, A, Saw, Saw, Saw, Q, Q, Q, Hammer, K, Key, Q, Saw, J, Key, K, Hammer, Q, Dynamite, A, A, A, Bat, J, Saw, Q, Hammer, K, Key, A, Hammer, J, Key, K, Bat, J, Dynamite, Q, Hammer, J, Hammer, A, A, Key, Q, Saw, K, Bat, J},
		// Reel 4
		{J, Bat, K, Saw, Q, Key, A, A, Hammer, J, Hammer, Q, Dynamite, J, Bat, K, Key, J, Hammer, A, Key, K, Hammer, Q, Saw, J, Bat, A, A, A, Dynamite, Q, Hammer, K, Key, J, Saw, Q, Key, K, Hammer, Q, Q, Q, Saw, Saw, Saw, A, Dynamite, K, Bat, Q, Saw, J, J, J, Hammer, A, A, A, Key, Key, Key, Q, Hammer, J, Saw, K, K, K, Bat, Bat, Bat, Q, Hammer, K, Key, J, Dynamite, Dynamite, Dynamite},
		// Reel 5
		{Dynamite, Dynamite, Dynamite, J, Key, K, Hammer, Q, Bat, Bat, Bat, K, K, K, Saw, J, Hammer, Q, Key, Key, Key, A, A, A, Hammer, J, J, J, Saw, Q, Bat, K, Dynamite, A, Saw, Saw, Saw, Q, Q, Q, Hammer, K, Key, Q, Saw, J, Key, K, Hammer, Q, Dynamite, A, A, A, Bat, J, Saw, Q, Hammer, K, Key, A, Hammer, J, Key, K, Bat, J, Dynamite, Q, Hammer, J, Hammer, A, A, Key, Q, Saw, K, Bat, J},
	},
}

// reel2 - Second reelset configuration with 93 symbols per reel and more SymbolBonus symbols
var reel2 = &Reels{
	Reels: [][]Symbol{
		// Reel 1
		{Dynamite, Dynamite, Dynamite, J, Key, K, Bonus, Hammer, Q, Bat, Bat, Bat, K, K, K, Bonus, Saw, J, Hammer, Q, Key, Key, Key, A, A, A, Hammer, Bonus, J, J, J, Saw, Q, Bat, K, Dynamite, A, Bonus, Bonus, Saw, Saw, Saw, Q, Q, Q, Hammer, K, Key, Q, Saw, J, Key, K, Bonus, Bonus, Bonus, Hammer, Q, Dynamite, A, A, A, Bat, J, Saw, Q, Hammer, Bonus, K, Key, A, Hammer, J, Key, K, Bat, J, Bonus, Bonus, Bonus, Dynamite, Q, Hammer, J, Hammer, A, A, Key, Q, Saw, K, Bat, J},
		// Reel 2
		{J, Bat, K, Saw, Q, Key, A, A, Hammer, J, Hammer, Q, Dynamite, Bonus, Bonus, Bonus, J, Bat, K, Key, J, Hammer, A, Key, K, Bonus, Hammer, Q, Saw, J, Bat, A, A, A, Dynamite, Q, Hammer, Bonus, Bonus, Bonus, K, Key, J, Saw, Q, Key, K, Hammer, Q, Q, Q, Saw, Saw, Saw, Bonus, Bonus, A, Dynamite, K, Bat, Q, Saw, J, J, J, Bonus, Hammer, A, A, A, Key, Key, Key, Q, Hammer, J, Saw, Bonus, K, K, K, Bat, Bat, Bat, Q, Hammer, Bonus, K, Key, J, Dynamite, Dynamite, Dynamite},
		// Reel 3
		{Dynamite, Dynamite, Dynamite, J, Key, K, Bonus, Hammer, Q, Bat, Bat, Bat, K, K, K, Bonus, Saw, J, Hammer, Q, Key, Key, Key, A, A, A, Hammer, Bonus, J, J, J, Saw, Q, Bat, K, Dynamite, A, Bonus, Bonus, Saw, Saw, Saw, Q, Q, Q, Hammer, K, Key, Q, Saw, J, Key, K, Bonus, Bonus, Bonus, Hammer, Q, Dynamite, A, A, A, Bat, J, Saw, Q, Hammer, Bonus, K, Key, A, Hammer, J, Key, K, Bat, J, Bonus, Bonus, Bonus, Dynamite, Q, Hammer, J, Hammer, A, A, Key, Q, Saw, K, Bat, J},
		// Reel 4
		{J, Bat, K, Saw, Q, Key, A, A, Hammer, J, Hammer, Q, Dynamite, Bonus, Bonus, Bonus, J, Bat, K, Key, J, Hammer, A, Key, K, Bonus, Hammer, Q, Saw, J, Bat, A, A, A, Dynamite, Q, Hammer, Bonus, Bonus, Bonus, K, Key, J, Saw, Q, Key, K, Hammer, Q, Q, Q, Saw, Saw, Saw, Bonus, Bonus, A, Dynamite, K, Bat, Q, Saw, J, J, J, Bonus, Hammer, A, A, A, Key, Key, Key, Q, Hammer, J, Saw, Bonus, K, K, K, Bat, Bat, Bat, Q, Hammer, Bonus, K, Key, J, Dynamite, Dynamite, Dynamite},
		// Reel 5
		{Dynamite, Dynamite, Dynamite, J, Key, K, Bonus, Hammer, Q, Bat, Bat, Bat, K, K, K, Bonus, Saw, J, Hammer, Q, Key, Key, Key, A, A, A, Hammer, Bonus, J, J, J, Saw, Q, Bat, K, Dynamite, A, Bonus, Bonus, Saw, Saw, Saw, Q, Q, Q, Hammer, K, Key, Q, Saw, J, Key, K, Bonus, Bonus, Bonus, Hammer, Q, Dynamite, A, A, A, Bat, J, Saw, Q, Hammer, Bonus, K, Key, A, Hammer, J, Key, K, Bat, J, Bonus, Bonus, Bonus, Dynamite, Q, Hammer, J, Hammer, A, A, Key, Q, Saw, K, Bat, J},
	},
}

// reel3 - Third reelset configuration with 81 symbols per reel (similar to reel1 but with different distribution)
var reel3 = &Reels{
	Reels: [][]Symbol{
		// Reel 1
		{Dynamite, Dynamite, Dynamite, J, Key, K, Hammer, Q, Bat, Bat, Bat, K, K, K, Saw, J, Hammer, Q, Key, Key, Key, A, A, A, Hammer, J, J, J, Saw, Q, Bat, K, Dynamite, A, Saw, Saw, Saw, Q, Q, Q, Hammer, K, Key, Q, Saw, J, Key, K, Hammer, Q, Dynamite, A, A, A, Bat, J, Saw, Q, Hammer, K, Key, A, Hammer, J, Key, K, Bat, J, Dynamite, Q, Hammer, J, Hammer, A, A, Key, Q, Saw, K, Bat, J},
		// Reel 2
		{J, Bat, K, Saw, Q, Key, A, A, Hammer, J, Hammer, Q, Dynamite, J, Bat, K, Key, J, Hammer, A, Key, K, Hammer, Q, Saw, J, Bat, A, A, A, Dynamite, Q, Hammer, K, Key, J, Saw, Q, Key, K, Hammer, Q, Q, Q, Saw, Saw, Saw, A, Dynamite, K, Bat, Q, Saw, J, J, J, Hammer, A, A, A, Key, Key, Key, Q, Hammer, J, Saw, K, K, K, Bat, Bat, Bat, Q, Hammer, K, Key, J, Dynamite, Dynamite, Dynamite},
		// Reel 3
		{Dynamite, Dynamite, Dynamite, J, Key, K, Hammer, Q, Bat, Bat, Bat, K, K, K, Saw, J, Hammer, Q, Key, Key, Key, A, A, A, Hammer, J, J, J, Saw, Q, Bat, K, Dynamite, A, Saw, Saw, Saw, Q, Q, Q, Hammer, K, Key, Q, Saw, J, Key, K, Hammer, Q, Dynamite, A, A, A, Bat, J, Saw, Q, Hammer, K, Key, A, Hammer, J, Key, K, Bat, J, Dynamite, Q, Hammer, J, Hammer, A, A, Key, Q, Saw, K, Bat, J},
		// Reel 4
		{J, Bat, K, Saw, Q, Key, A, A, Hammer, J, Hammer, Q, Dynamite, J, Bat, K, Key, J, Hammer, A, Key, K, Hammer, Q, Saw, J, Bat, A, A, A, Dynamite, Q, Hammer, K, Key, J, Saw, Q, Key, K, Hammer, Q, Q, Q, Saw, Saw, Saw, A, Dynamite, K, Bat, Q, Saw, J, J, J, Hammer, A, A, A, Key, Key, Key, Q, Hammer, J, Saw, K, K, K, Bat, Bat, Bat, Q, Hammer, K, Key, J, Dynamite, Dynamite, Dynamite},
		// Reel 5
		{Dynamite, Dynamite, Dynamite, J, Key, K, Hammer, Q, Bat, Bat, Bat, K, K, K, Saw, J, Hammer, Q, Key, Key, Key, A, A, A, Hammer, J, J, J, Saw, Q, Bat, K, Dynamite, A, Saw, Saw, Saw, Q, Q, Q, Hammer, K, Key, Q, Saw, J, Key, K, Hammer, Q, Dynamite, A, A, A, Bat, J, Saw, Q, Hammer, K, Key, A, Hammer, J, Key, K, Bat, J, Dynamite, Q, Hammer, J, Hammer, A, A, Key, Q, Saw, K, Bat, J},
	},
}

// reel4 - Fourth reelset configuration with 96 symbols per reel with many SymbolBonus symbols
var reel4 = &Reels{
	Reels: [][]Symbol{
		// Reel 1
		{Dynamite, Dynamite, Dynamite, J, Key, K, Bonus, Hammer, Q, Bat, Bat, Bat, K, K, K, Bonus, Saw, J, Hammer, Q, Key, Key, Key, A, A, A, Hammer, Bonus, J, J, J, Saw, Q, Bat, K, Dynamite, A, Bonus, Bonus, Saw, Saw, Saw, Q, Q, Q, Hammer, K, Bonus, Bonus, Bonus, Key, Q, Saw, J, Key, K, Bonus, Bonus, Bonus, Hammer, Q, Dynamite, A, A, A, Bat, J, Saw, Q, Hammer, Bonus, K, Key, A, Hammer, J, Key, K, Bat, J, Bonus, Bonus, Bonus, Dynamite, Q, Hammer, J, Hammer, A, A, Key, Q, Saw, K, Bat, J},
		// Reel 2
		{J, Bat, K, Saw, Q, Key, A, A, Hammer, J, Hammer, Q, Dynamite, Bonus, Bonus, Bonus, J, Bat, K, Key, J, Hammer, A, Key, K, Bonus, Hammer, Q, Saw, J, Bat, A, A, A, Dynamite, Q, Hammer, Bonus, Bonus, Bonus, K, Key, J, Saw, Q, Key, K, Bonus, Bonus, Bonus, Hammer, Q, Q, Q, Saw, Saw, Saw, Bonus, Bonus, A, Dynamite, K, Bat, Q, Saw, J, J, J, Bonus, Hammer, A, A, A, Key, Key, Key, Q, Hammer, J, Saw, Bonus, K, K, K, Bat, Bat, Bat, Q, Hammer, Bonus, K, Key, J, Dynamite, Dynamite, Dynamite},
		// Reel 3
		{Dynamite, Dynamite, Dynamite, J, Key, K, Bonus, Hammer, Q, Bat, Bat, Bat, K, K, K, Bonus, Saw, J, Hammer, Q, Key, Key, Key, A, A, A, Hammer, Bonus, J, J, J, Saw, Q, Bat, K, Dynamite, A, Bonus, Bonus, Saw, Saw, Saw, Q, Q, Q, Hammer, K, Bonus, Bonus, Bonus, Key, Q, Saw, J, Key, K, Bonus, Bonus, Bonus, Hammer, Q, Dynamite, A, A, A, Bat, J, Saw, Q, Hammer, Bonus, K, Key, A, Hammer, J, Key, K, Bat, J, Bonus, Bonus, Bonus, Dynamite, Q, Hammer, J, Hammer, A, A, Key, Q, Saw, K, Bat, J},
		// Reel 4
		{J, Bat, K, Saw, Q, Key, A, A, Hammer, J, Hammer, Q, Dynamite, Bonus, Bonus, Bonus, J, Bat, K, Key, J, Hammer, A, Key, K, Bonus, Hammer, Q, Saw, J, Bat, A, A, A, Dynamite, Q, Hammer, Bonus, Bonus, Bonus, K, Key, J, Saw, Q, Key, K, Bonus, Bonus, Bonus, Hammer, Q, Q, Q, Saw, Saw, Saw, Bonus, Bonus, A, Dynamite, K, Bat, Q, Saw, J, J, J, Bonus, Hammer, A, A, A, Key, Key, Key, Q, Hammer, J, Saw, Bonus, K, K, K, Bat, Bat, Bat, Q, Hammer, Bonus, K, Key, J, Dynamite, Dynamite, Dynamite},
		// Reel 5
		{Dynamite, Dynamite, Dynamite, J, Key, K, Bonus, Hammer, Q, Bat, Bat, Bat, K, K, K, Bonus, Saw, J, Hammer, Q, Key, Key, Key, A, A, A, Hammer, Bonus, J, J, J, Saw, Q, Bat, K, Dynamite, A, Bonus, Bonus, Saw, Saw, Saw, Q, Q, Q, Hammer, K, Bonus, Bonus, Bonus, Key, Q, Saw, J, Key, K, Bonus, Bonus, Bonus, Hammer, Q, Dynamite, A, A, A, Bat, J, Saw, Q, Hammer, Bonus, K, Key, A, Hammer, J, Key, K, Bat, J, Bonus, Bonus, Bonus, Dynamite, Q, Hammer, J, Hammer, A, A, Key, Q, Saw, K, Bat, J},
	},
}
