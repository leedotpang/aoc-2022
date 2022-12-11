package main

import (
	"fmt"
	"lee/aoc/helpers"
	"os"
	"strconv"
	"strings"
)

type CommandQueue = helpers.Stack[int]

func find_solution(content string) int {
	targets := map[int]bool{20: true, 60: true, 100: true, 140: true, 180: true, 220: true}
	register, sum_of_signals, value := 1, 0, 0
	operation := []string{}
	command := ""
	// pad queue by 1 for cycles to match
	cmd_queue := CommandQueue{0}

	// Loop through ops to establish proper queue
	for _, op_string := range strings.Split(content, "\n") {
		operation = strings.Split(op_string, " ")
		command = operation[0]
		cmd_queue.Push(0)

		if command == "addx" {
			value, _ = strconv.Atoi(operation[1])
			cmd_queue.Push(value)
		}
	}

	// Is it possible to do this all at once?
	for cycle, value := range cmd_queue {
		if targets[cycle] {
			sum_of_signals += register * cycle
		}

		register += value
	}
	return sum_of_signals
}

func main() {
	// read input
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print(fmt.Errorf("%w", err))
	}
	content_str := string(content)
	// 	content_str := `addx 15
	// addx -11
	// addx 6
	// addx -3
	// addx 5
	// addx -1
	// addx -8
	// addx 13
	// addx 4
	// noop
	// addx -1
	// addx 5
	// addx -1
	// addx 5
	// addx -1
	// addx 5
	// addx -1
	// addx 5
	// addx -1
	// addx -35
	// addx 1
	// addx 24
	// addx -19
	// addx 1
	// addx 16
	// addx -11
	// noop
	// noop
	// addx 21
	// addx -15
	// noop
	// noop
	// addx -3
	// addx 9
	// addx 1
	// addx -3
	// addx 8
	// addx 1
	// addx 5
	// noop
	// noop
	// noop
	// noop
	// noop
	// addx -36
	// noop
	// addx 1
	// addx 7
	// noop
	// noop
	// noop
	// addx 2
	// addx 6
	// noop
	// noop
	// noop
	// noop
	// noop
	// addx 1
	// noop
	// noop
	// addx 7
	// addx 1
	// noop
	// addx -13
	// addx 13
	// addx 7
	// noop
	// addx 1
	// addx -33
	// noop
	// noop
	// noop
	// addx 2
	// noop
	// noop
	// noop
	// addx 8
	// noop
	// addx -1
	// addx 2
	// addx 1
	// noop
	// addx 17
	// addx -9
	// addx 1
	// addx 1
	// addx -3
	// addx 11
	// noop
	// noop
	// addx 1
	// noop
	// addx 1
	// noop
	// noop
	// addx -13
	// addx -19
	// addx 1
	// addx 3
	// addx 26
	// addx -30
	// addx 12
	// addx -1
	// addx 3
	// addx 1
	// noop
	// noop
	// noop
	// addx -9
	// addx 18
	// addx 1
	// addx 2
	// noop
	// noop
	// addx 9
	// noop
	// noop
	// noop
	// addx -1
	// addx 2
	// addx -37
	// addx 1
	// addx 3
	// noop
	// addx 15
	// addx -21
	// addx 22
	// addx -6
	// addx 1
	// noop
	// addx 2
	// addx 1
	// noop
	// addx -10
	// noop
	// noop
	// addx 20
	// addx 1
	// addx 2
	// addx 2
	// addx -6
	// addx -11
	// noop
	// noop
	// noop`

	fmt.Println("Sum of signal strengths: ", find_solution(content_str))
}
