package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func find_solution(content string) int {
	string_groups := strings.Split(content, "\n\n")
	totals := []int{}
	top_three_total := 0

	for _, group := range string_groups {
		curr_total := 0
		for _, num2 := range strings.Split(group, "\n") {
			curr_int, _ := strconv.Atoi(num2)
			curr_total += curr_int
		}
		totals = append(totals, curr_total)
	}

	sort.Slice(totals, func(i, j int) bool {
		return totals[i] > totals[j]
	})

	for _, num := range totals[:3] {
		top_three_total += num
	}

	return top_three_total
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Errorf("%w", err)
	}
	content_str := string(content)
	fmt.Println("Top Three Total Calories: ", find_solution(content_str))
}
