package main

import (
	"fmt"
	"os"
	"strings"
)

func already_counted(i int, length int) bool {
	return i == 0 || i == length-1
}

func visible_vertically(tree int, hi int, vi int, rows []string) bool {
	visible_down, visible_up := true, true
	for index, tree_compare := range rows {
		if int(tree_compare[hi]-'0') >= tree {
			switch {
			case index < vi:
				visible_up = false
			case index > vi:
				visible_down = false
			}
		}
	}
	return visible_up || visible_down
}

func visible_horizontally(tree int, hi int, row string) bool {
	visible_left, visible_right := true, true
	for index, tree_compare := range row {
		if int(tree_compare-'0') >= tree {
			switch {
			case index < hi:
				visible_left = false
			case index > hi:
				visible_right = false
			}
		}
	}
	return visible_left || visible_right
}

func tree_is_visible(tree rune, vi int, hi int, rows []string) bool {
	tree_val := int(tree - '0')
	// tree is visible if ANY
	// visible from top, bottom, left, or right
	// from top = same index until no more rows--
	// from bottom = same index until no more rows++
	// from left = same string--
	// from right = same string ++
	return visible_vertically(tree_val, hi, vi, rows) || visible_horizontally(tree_val, hi, rows[vi])
}

func find_solution(content string) int {
	visible_trees := 0
	tree_rows := strings.Split(content, "\n")
	tree_rows_len := len(tree_rows)
	// Count perimiter in characters
	// Add row length
	visible_trees += len(tree_rows) * 2
	// Add col length minus overlapping row trees
	visible_trees += (len(tree_rows[0]) - 2) * 2

	// count numbers in grid that are bigger than all adjacent numbers to the edge
	for vi, tree_row := range tree_rows {
		if already_counted(vi, tree_rows_len) {
			continue
		}
		for hi, tree := range tree_row {
			if already_counted(hi, len(tree_row)) {
				continue
			}
			if tree_is_visible(tree, vi, hi, tree_rows) {
				visible_trees++
			}
		}
	}
	return visible_trees
}

func main() {
	// read input
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print(fmt.Errorf("%w", err))
	}
	content_str := string(content)
	// 	content_str := `30373
	// 25512
	// 65332
	// 33549
	// 35390`

	fmt.Println("# of Visible Trees: ", find_solution(content_str))
}
