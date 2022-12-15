package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Im trying to catch up at this point because the last few days have me behind, so instead of creating a parser, I'm going to live dangerously and use the built in parser to read the arrays and run the comparisons
// Also, I'm slowly learning best practice, so my files will start to look different as I continue to see better-formatted examples
func main() {
	// read input
	// start := time.Now()
	content, _ := os.ReadFile("input.txt")
	content_str := string(content)

	fmt.Println("Sum of indices: ", find_solution(content_str))
	// fmt.Println(time.Since(start).Seconds())
}

func find_solution(content string) int {
	// Parse input into lists of ints
	pairs := strings.Split(content, "\n\n")
	sum_of_indices := 0
	curr_result := 0

	for index, raw_pair := range pairs {
		pair := strings.Split(raw_pair, "\n")
		var first, second []any
		json.Unmarshal([]byte(pair[0]), &first)
		json.Unmarshal([]byte(pair[1]), &second)

		curr_result = compare_lists(first, second)
		if curr_result < 0 {
			sum_of_indices += index + 1
		}
	}

	return sum_of_indices
}

// Since I don't have as much time to get these questions 100% solved -- I'm going through until I'm stuck enough to find a better solution elsewhere to implement in my own way. Often I'll find a solution in a different language to at least add the challenge of translating it to Go. (this beautiful solution was stolen <3)
func compare_lists(p1, p2 any) int {

	a, a_ok := p1.([]any)
	b, b_ok := p2.([]any)

	switch {
	case !a_ok && !b_ok:
		return int(p1.(float64) - p2.(float64))
	case !a_ok:
		a = []any{p1}
	case !b_ok:
		b = []any{p2}
	}

	for i := 0; i < len(a) && i < len(b); i++ {
		if c := compare_lists(a[i], b[i]); c != 0 {
			return c
		}
	}

	// if indexes are both lists, left must have fewer children
	return len(a) - len(b)
}
