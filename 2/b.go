package main

import (
	"fmt"
	"os"
	"strings"
)

func find_solution(content string) int {
	// A = Rock 1
	// B = Paper 2
	// C = Scissors 3
	// Win - Z - 6 / Lose - X -  0 / Draw - Y - 3
	rps_key := map[string]int{
		"AX": 3,
		"BX": 1,
		"CX": 2,
		"AY": 4,
		"BY": 5,
		"CY": 6,
		"AZ": 8,
		"BZ": 9,
		"CZ": 7,
	}

	total_points := 0

	for _, raw_game := range strings.Split(content, "\n") {
		game := strings.Replace(raw_game, " ", "", 1)
		total_points += rps_key[game]
	}

	return total_points
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print(fmt.Errorf("%w", err))
	}
	content_str := string(content)
	fmt.Println("Total RPS Score: ", find_solution(content_str))
}
