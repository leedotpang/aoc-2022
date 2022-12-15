package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	content, _ := os.ReadFile("input.txt")
	content_str := string(content)

	fmt.Println("Sum of indices: ", find_solution(content_str))
}

func find_solution(content string) int {
	pairs := strings.Split(content, "\n\n")
	pairs = append(pairs, "[[2]]\n[[6]]")
	dividers_multiplied := 1
	ordered_pairs := []any{}

	for _, raw_pair := range pairs {
		pair := strings.Split(raw_pair, "\n")
		var first, second []any
		json.Unmarshal([]byte(pair[0]), &first)
		json.Unmarshal([]byte(pair[1]), &second)

		ordered_pairs = append(ordered_pairs, []any{first, second}...)
	}

	// sort ordered pairs
	sort.Slice(ordered_pairs, func(i, j int) bool { return compare_lists(ordered_pairs[i], ordered_pairs[j]) < 0 })

	// find index of dividers and multiply them for the result
	for index, pair := range ordered_pairs {
		if fmt.Sprint(pair) == "[[2]]" || fmt.Sprint(pair) == "[[6]]" {
			dividers_multiplied *= index + 1
		}
	}

	return dividers_multiplied
}

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
