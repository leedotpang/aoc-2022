package main

import (
	"fmt"
	"lee/aoc/helpers"
	"math"
	"os"
	"strconv"
	"strings"
)

type CommandQueue = helpers.Stack[int]

func find_solution(content string) {
	register, value, diff := 1, 0, 0
	operation := []string{}
	command := ""
	screen_row := ""
	cmd_queue := CommandQueue{}

	for _, op_string := range strings.Split(content, "\n") {
		operation = strings.Split(op_string, " ")
		command = operation[0]
		cmd_queue.Push(0)

		if command == "addx" {
			value, _ = strconv.Atoi(operation[1])
			cmd_queue.Push(value)
		}
	}

	for cycle, value := range cmd_queue {
		diff = int(math.Abs(float64(register - cycle%40)))

		if diff == 1 || diff == 0 {
			screen_row += "#"
		} else {
			screen_row += "."
		}

		if len(screen_row) == 40 {
			fmt.Println(screen_row)
			screen_row = ""
		}

		register += value
	}
}

func main() {
	// read input
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print(fmt.Errorf("%w", err))
	}
	content_str := string(content)
	find_solution(content_str)
}
