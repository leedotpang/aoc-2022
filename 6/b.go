package main

import (
	"fmt"
	"os"
)

const MARKER_TEST_LENGTH = 14

func unique_string(s string) bool {
	checker := map[rune]bool{}
	for _, char := range s {
		_, exists := checker[char]
		if exists {
			return false
		}
		checker[char] = true
	}
	return true
}

func find_solution(content string) int {
	for position := range []rune(content) {
		test := content[position : position+MARKER_TEST_LENGTH]
		if unique_string(test) {
			return position + MARKER_TEST_LENGTH
		}
	}
	return 0
}

func main() {
	// read input
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print(fmt.Errorf("%w", err))
	}
	content_str := string(content)

	fmt.Println("First message marker position: ", find_solution(content_str))
}
