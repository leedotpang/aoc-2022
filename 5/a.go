package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Stack []string

func (s *Stack) Push(el string) bool {
	*s = append(*s, el)
	return true
}

func (s *Stack) Pop() string {
	last_index := len(*s) - 1
	last := (*s)[last_index]
	*s = (*s)[:last_index]
	return last
}

func (s *Stack) Reverse() {
	var reversed_s Stack
	for i := len(*s) - 1; i >= 0; i-- {
		reversed_s.Push((*s)[i])
	}
	*s = reversed_s
}

func chunk_string(s string, size int) []string {
	var chunks []string
	length := len(s)
	runes := []rune(s)

	if length == 0 {
		return []string{s}
	}

	for i := 0; i < length; i += size {
		nn := i + size
		if nn > length {
			nn = length
		}
		chunks = append(chunks, string(runes[i:nn]))
	}
	return chunks
}

func prepend_crate(row []string, crate string) []string {
	row = append(row, "")
	row[0] = crate

	return row
}

func process_stack_rearrangement(stacks []Stack, count int, from int, to int) {
	// take COUNT from FROM to TO
	for i := 0; i < count; i++ {
		take := stacks[from].Pop()
		stacks[to].Push(take)
	}
}

func rows_to_col_stacks(stacks [][]string) []Stack {
	new_stack := []Stack{}
	col_count := len(stacks[len(stacks)-1])

	for i := 0; i < col_count; i++ {
		var temp_stack Stack
		for _, stack := range stacks {
			stack_len := len(stack)
			in_range := stack_len > i

			if !in_range {
				continue
			}

			if val := strings.Trim(stack[i], " "); val != "" {
				temp_stack.Push(val)
			}
		}
		temp_stack.Reverse()
		new_stack = append(new_stack, temp_stack)
	}

	fmt.Println(new_stack, col_count)

	return new_stack
}

func get_tops_of_stacks(stacks []Stack) string {
	tops_of_stacks := ""
	for _, stack := range stacks {
		tops_of_stacks += stack.Pop()
	}
	tops_of_stacks = regexp.MustCompile(`\W*`).ReplaceAllString(tops_of_stacks, "")
	return tops_of_stacks
}

func find_solution(content string) string {
	// split input between crates and instructions
	input_parts := strings.Split(content, "\n\n")
	stack_rows := input_parts[0]
	rearrangements := input_parts[1]
	raw_stack_data := [][]string{}
	top_of_each_stack := ""

	// Loop over lines and chunk into cols by 4 spaces (this is fragile). Then remove last row (col numbers) and reassign to the stacks in col-based slices.
	for _, row := range strings.Split(stack_rows, "\n") {
		temp_row := []string{}
		for _, crate := range chunk_string(row, 4) {
			temp_row = append(temp_row, crate)
		}
		raw_stack_data = append(raw_stack_data, temp_row)
	}
	raw_stack_data = raw_stack_data[:len(raw_stack_data)-1]

	stacks := rows_to_col_stacks(raw_stack_data)

	for _, procedure := range strings.Split(rearrangements, "\n") {
		if len(procedure) < 1 {
			continue
		}
		nums := regexp.MustCompile(`\d+`).FindAllString(procedure, 3)
		fmt.Println(nums, procedure)
		count, _ := strconv.Atoi(nums[0])
		from, _ := strconv.Atoi(nums[1])
		to, _ := strconv.Atoi(nums[2])
		process_stack_rearrangement(stacks, count, from-1, to-1)
	}

	fmt.Println()

	top_of_each_stack = get_tops_of_stacks(stacks)
	return top_of_each_stack
}

func main() {
	// read input
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print(fmt.Errorf("%w", err))
	}
	content_str := string(content)
	// content_str := "    [D]     \n" +
	// 	"[N] [C]     \n" +
	// 	"[Z] [M] [P]    \n" +
	// 	" 1  2   3   \n\n" +
	// 	`move 1 from 2 to 1
	// move 3 from 1 to 3
	// move 2 from 2 to 1
	// move 1 from 1 to 2`

	fmt.Println("Top of Stacks: ", find_solution(content_str))
}
