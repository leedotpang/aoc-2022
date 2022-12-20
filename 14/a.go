package main

import (
	"fmt"
	"image"
	"os"
	"strings"
)

func main() {
	content, _ := os.ReadFile("input.txt")
	content_str := string(content)

	fmt.Println("Units of sand at rest: ", find_solution(content_str))
}

func find_solution(content string) int {
	var x, y int
	blocked := map[image.Point]bool{}
	last_position, curr_position := image.Point{}, image.Point{}
	movements := []image.Point{
		{0, 1},  // Down
		{-1, 1}, // Left and down
		{1, 1},  // Right and down
	}
	total_units := 0

	// parse lines for blocked squares
	// add inaccesible squares in a slice
	for _, line := range strings.Split(content, "\n") {
		last_position = image.Point{-999, -999}
		for _, raw_position := range strings.Split(line, " -> ") {
			fmt.Sscanf(raw_position, "%d,%d", &x, &y)

			curr_position = image.Point{x, y}
			blocked[image.Point{x, y}] = true

			diffX := last_position.X-curr_position.X != 0
			diffY := last_position.Y-curr_position.Y != 0

			switch {
			case last_position.X == -999:
				// do nothing for this step
			case diffX:
				// loop through and fill
				hi := map[bool]int{true: last_position.X, false: curr_position.X}[last_position.X > curr_position.X]
				lo := map[bool]int{true: curr_position.X, false: last_position.X}[last_position.X > curr_position.X]
				for i := lo; i < hi; i++ {
					blocked[image.Point{i, y}] = true
				}
			case diffY:
				// ^^ same
				hi := map[bool]int{true: last_position.Y, false: curr_position.Y}[last_position.Y > curr_position.Y]
				lo := map[bool]int{true: curr_position.Y, false: last_position.Y}[last_position.Y > curr_position.Y]
				for i := lo; i < hi; i++ {
					blocked[image.Point{x, i}] = true
				}
			}

			last_position = curr_position
		}
	}

	// set max height of area sand can fall
	max_height := 0
	for p := range blocked {
		if p.Y > max_height {
			max_height = p.Y
		}
	}

	// squares lead either one down and left, or one down and right --> or it stays where it lands if it cant move down further

	// simulate falling sand from 500,0
	for {
		curr_position = image.Point{500, 0}

		// Loop through all possible movements and do them until sand is blocked
		for i := 0; i < 3; i++ {
			move := movements[i]

			if _, stop := blocked[curr_position.Add(move)]; !stop && curr_position.Add(move).Y <= max_height {
				curr_position = curr_position.Add(move)
				i = -1
			}
		}

		if curr_position.Y == max_height {
			break
		}
		blocked[curr_position] = true
		total_units++
	}

	return total_units
}
