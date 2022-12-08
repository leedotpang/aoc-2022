package main

import (
	"fmt"
	"os"
	"strings"
)

type Match []int

func (m *Match) Push(i int) {
	(*m) = append(*m, i)
}

func (m *Match) Pop() int {
	return (*m)[len(*m)-1]
}

func (m *Match) Shift() int {
	return (*m)[0]
}

func (m *Match) Present() bool {
	return len(*m) > 0
}

func get_top_score(tree int, hi int, vi int, rows []string) int {
	matches := Match{}

	for index, tree_compare := range rows {
		if index == vi {
			break
		}
		if int(tree_compare[hi]-'0') >= tree {
			matches.Push(vi - index)
		}
	}

	if matches.Present() {
		return matches.Pop()
	}

	return vi
}

func get_bottom_score(tree int, hi int, vi int, rows []string) int {
	matches := Match{}

	for index, tree_compare := range rows {
		if index <= vi {
			continue
		}
		if int(tree_compare[hi]-'0') >= tree {
			matches.Push(index - vi)
		}
	}

	if matches.Present() {
		return matches.Shift()
	}

	return len(rows) - 1 - vi
}

func get_left_score(tree int, hi int, row string) int {
	matches := Match{}
	for index, tree_compare := range row {
		if index == hi {
			break
		}
		if int(tree_compare-'0') >= tree {
			matches.Push(hi - index)
		}
	}

	if matches.Present() {
		return matches.Pop()
	}

	return hi
}

func get_right_score(tree int, hi int, row string) int {
	matches := Match{}

	for index, tree_compare := range row {
		if index <= hi {
			continue
		}
		if int(tree_compare-'0') >= tree {
			matches.Push(index - hi)
		}
	}

	if matches.Present() {
		return matches.Shift()
	}

	return len(row) - 1 - hi
}

func get_scenic_score(tree rune, hi int, vi int, rows []string) int {
	tree_val := int(tree - '0')
	top := get_top_score(tree_val, hi, vi, rows)
	bottom := get_bottom_score(tree_val, hi, vi, rows)
	left := get_left_score(tree_val, hi, rows[vi])
	right := get_right_score(tree_val, hi, rows[vi])

	// loop horizontally/vertically
	// ** select CLOSEST matching element for calculation
	// multiply # of trees visible around tree up to a tree that blocks it
	return top * bottom * left * right
}

func find_solution(content string) int {
	tree_rows := strings.Split(content, "\n")
	highest_score, curr_score := 0, 0

	for vi, tree_row := range tree_rows {
		for hi, tree := range tree_row {
			curr_score = get_scenic_score(tree, hi, vi, tree_rows)

			if curr_score > highest_score {
				highest_score = curr_score
			}
		}
	}
	return highest_score
}

func main() {
	// read input
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print(fmt.Errorf("%w", err))
	}
	content_str := string(content)

	fmt.Println("Highest Scenic Score: ", find_solution(content_str))
}
