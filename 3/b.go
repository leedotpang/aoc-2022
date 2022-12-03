package main

import (
	"fmt"
	"os"
	"strings"
)

const ALPHABET string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func alphabetIndex(char rune) int {
	return strings.IndexRune(ALPHABET, char) + 1
}

func find_solution(content string) int {
	total_points := 0
	rucksacks := strings.Split(content, "\n")

	for index, sack_one := range rucksacks {
		if index%3 != 0 {
			continue
		}
		sack_two := rucksacks[index+1]
		sack_three := rucksacks[index+2]

		for _, char := range sack_one {
			match := strings.IndexRune(sack_two, char) > -1 && strings.IndexRune(sack_three, char) > -1
			if match {
				total_points += alphabetIndex(char)
				break
			}
		}
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
	fmt.Println("Sum of Priorities: ", find_solution(content_str))
}
