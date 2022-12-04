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
	overlapped_pairs := 0

	for _, elf_pair := range strings.Split(content, "\n") {
		// switched to map for easier matching between arrays/slices
		full_assignments := [2]map[string]bool{{}, {}}

		for index, raw_assignment := range strings.Split(elf_pair, ",") {
			start, end := split_assignment(raw_assignment)

			for i := start; i <= end; i++ {
				full_assignments[index][strconv.Itoa(i)] = true
			}

		}

		for section, _ := range full_assignments[0] {
			if _, overlapped := full_assignments[1][section]; overlapped {
				overlapped_pairs++
				break
			}
		}

	}

	return overlapped_pairs
}

func main() {
	// read input
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print(fmt.Errorf("%w", err))
	}
	content_str := string(content)

	fmt.Println("# of overlapped groups: ", find_solution(content_str))
}
