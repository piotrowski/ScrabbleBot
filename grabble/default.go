package grabble

// CreateDefaultGame - it creates and return basic Grabble game.
func CreateDefaultGame(players []string) Grabble {
	officialScrabbleBoard := [15][15]rune{
		{'W', '0', '0', 'l', '0', '0', '0', 'W', '0', '0', '0', 'l', '0', '0', 'W'},
		{'0', 'w', '0', '0', '0', 'L', '0', '0', '0', 'L', '0', '0', '0', 'w', '0'},
		{'0', '0', 'w', '0', '0', '0', 'l', '0', 'l', '0', '0', '0', 'w', '0', '0'}, // 3
		{'l', '0', '0', 'w', '0', '0', '0', 'l', '0', '0', '0', 'w', '0', '0', 'l'},
		{'0', '0', '0', '0', 'w', '0', '0', '0', '0', '0', 'w', '0', '0', '0', '0'},
		{'0', 'L', '0', '0', '0', 'L', '0', '0', '0', 'L', '0', '0', '0', 'L', '0'}, // 6
		{'0', '0', 'l', '0', '0', '0', 'l', '0', 'l', '0', '0', '0', 'l', '0', '0'},
		{'W', '0', '0', 'l', '0', '0', '0', 's', '0', '0', '0', 'l', '0', '0', 'W'},
		{'0', '0', 'l', '0', '0', '0', 'l', '0', 'l', '0', '0', '0', 'l', '0', '0'}, // 9
		{'0', 'L', '0', '0', '0', 'L', '0', '0', '0', 'L', '0', '0', '0', 'L', '0'},
		{'0', '0', '0', '0', 'w', '0', '0', '0', '0', '0', 'w', '0', '0', '0', '0'},
		{'l', '0', '0', 'w', '0', '0', '0', 'l', '0', '0', '0', 'w', '0', '0', 'l'}, // 12
		{'0', '0', 'w', '0', '0', '0', 'l', '0', 'l', '0', '0', '0', 'w', '0', '0'},
		{'0', 'w', '0', '0', '0', 'L', '0', '0', '0', 'L', '0', '0', '0', 'w', '0'},
		{'W', '0', '0', 'l', '0', '0', '0', 'W', '0', '0', '0', 'l', '0', '0', 'W'}, // 15
	}
	officialDict := "exampleData/collins_official_scrabble_2019.txt"
	tilePoints := map[rune]int{
		'_': 0,
		'E': 1, 'A': 1, 'I': 1, 'O': 1, 'N': 1, 'R': 1, 'T': 1, 'L': 1, 'S': 1, 'U': 1,
		'D': 2, 'G': 2,
		'B': 3, 'C': 3, 'M': 3, 'P': 3,
		'F': 4, 'H': 4, 'V': 4, 'W': 4, 'Y': 4,
		'K': 5,
		'J': 8, 'X': 8,
		'Q': 10, 'Z': 10,
	}
	allTiles := []rune("__EEEEEEEEEEEEAAAAAAAAAIIIIIIIIIOOOOOOOONNNNNNRRRRRRTTTTTTLLLLSSSSUUUUDDDDGGGBBCCMMPPFFHHVVWWYYKJXQZ")

	return CreateGrabble(officialDict, officialScrabbleBoard, players, allTiles, tilePoints, 7)
}