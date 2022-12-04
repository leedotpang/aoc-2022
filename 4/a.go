package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func split_assignment(raw_assignment string) (int, int) {
	assignment := strings.SplitN(raw_assignment, "-", 2)
	start, _ := strconv.Atoi(assignment[0])
	end, _ := strconv.Atoi(assignment[1])
	return start, end
}

func find_solution(content string) int {
	fully_overlapped_pairs := 0

	for _, elf_pair := range strings.Split(content, "\n") {
		full_assignments := [2]string{}

		// split groups by ,
		for index, raw_assignment := range strings.Split(elf_pair, ",") {
			// loop over each group and fill numbers in between
			start, end := split_assignment(raw_assignment)

			for i := start; i <= end; i++ {
				full_assignments[index] += " " + strconv.Itoa(i) + " "
				// add separators to distinguish multidigit numbers for string compare
			}

		}

		overlapped := strings.Contains(full_assignments[0], full_assignments[1]) || strings.Contains(full_assignments[1], full_assignments[0])

		if overlapped {
			fully_overlapped_pairs++
		}
	}

	return fully_overlapped_pairs
}

func main() {
	// read input
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print(fmt.Errorf("%w", err))
	}
	content_str := string(content)

	fmt.Println("# of fully overlapped groups: ", find_solution(content_str))
}
