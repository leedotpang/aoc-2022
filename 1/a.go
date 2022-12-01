package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func find_solution(content string) int {
	// split input by blank lines to get groupings
	string_groups := strings.Split(content, "\n\n")

	// total each group
	biggest_total := 0

	for _, group := range string_groups {
		curr_total := 0
		for _, num2 := range strings.Split(group, "\n") {
			curr_int, _ := strconv.Atoi(num2)
			curr_total += curr_int
		}

		if curr_total > biggest_total {
			biggest_total = curr_total
		}
	}

	// what is the biggest total?
	return biggest_total
}

func main() {
	// read input
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Errorf("%w", err)
	}
	content_str := string(content)
	fmt.Println("Most calories: ", find_solution(content_str))
}
