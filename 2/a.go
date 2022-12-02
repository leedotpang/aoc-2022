package main

import (
	"fmt"
	"os"
	"strings"
)

func find_solution(content string) int {
	// create map of values for each move
	// A/X = Rock 1
	// B/Y = Paper 2
	// C/Z = Scissors 3
	// Win 6 / Lose 0 / Draw 3
	rps_key := map[string]int{
		"AX": 4,
		"BX": 1,
		"CX": 7,
		"AY": 8,
		"BY": 5,
		"CY": 2,
		"AZ": 3,
		"BZ": 9,
		"CZ": 6,
	}

	total_points := 0

	// loop over games and calculate running total for points
	for _, raw_game := range strings.Split(content, "\n") {
		game := strings.Replace(raw_game, " ", "", 1)
		total_points += rps_key[game]
	}

	return total_points
}

func main() {
	// read input
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print(fmt.Errorf("%w", err))
	}
	content_str := string(content)
	fmt.Println("Total RPS Score: ", find_solution(content_str))
}
