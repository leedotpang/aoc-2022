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

	// loop over lines
	for _, rucksack := range strings.Split(content, "\n") {
		// split line into 2 compartments
		compartment_size := len(rucksack) / 2
		compartment_one := rucksack[:compartment_size]
		compartment_two := rucksack[compartment_size:]

		// find matching letter in each
		for _, char := range compartment_one {
			if strings.IndexRune(compartment_two, char) > -1 {
				// add int value of that character into running total
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
