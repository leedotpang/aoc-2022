package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Over complicating this with types to actually learn :)
type Direction int

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
)

func (d Direction) String() string {
	switch d {
	case UP:
		return "U"
	case DOWN:
		return "D"
	case LEFT:
		return "L"
	case RIGHT:
		return "R"
	default:
		return "undefined"
	}
}

func toDirection(s string) Direction {
	dirs := []Direction{UP, DOWN, LEFT, RIGHT}

	for _, dir := range dirs {
		if dir.String() == s {
			return dir
		}
	}

	// Defaults to UP for reasons
	return UP
}

type Coordinate struct {
	x int
	y int
}

func (p *Coordinate) Move(dir Direction) {
	switch dir {
	case UP:
		p.y += 1
	case DOWN:
		p.y -= 1
	case LEFT:
		p.x -= 1
	case RIGHT:
		p.x += 1
	}
}

func catch_up(head Coordinate, tail *Coordinate, visited *map[Coordinate]bool) {
	// Don't move tail if head is adjacent
	// if head moves past adjacent on a different axis, shift tail behind head and onto the correct axis -- this is the diaganol movement
	var dir Direction
	x_diff := int(math.Abs(float64(head.x - tail.x)))
	y_diff := int(math.Abs(float64(head.y - tail.y)))

	// Adjacent == x diff < 2 && y diff < 2
	if x_diff < 2 && y_diff < 2 {
		return
	}
	if x_diff != 0 {
		// I'm not sure if you're supposed to do this in Go but it reminds me of a ternary and makes me happy
		dir = map[bool]Direction{true: RIGHT, false: LEFT}[head.x-tail.x > 0]
		tail.Move(dir)
	}
	if y_diff != 0 {
		dir = map[bool]Direction{true: UP, false: DOWN}[head.y-tail.y > 0]
		tail.Move(dir)
	}
	(*visited)[*tail] = true
}

func find_solution(content string) int {
	head, tail := &Coordinate{x: 0, y: 0}, &Coordinate{x: 0, y: 0}
	instructions := strings.Split(content, "\n")
	moves := 0
	currInstructionSet := []string{}
	var direction Direction
	visited := map[Coordinate]bool{*tail: true}

	// Loop over instructions and move in required direction one move at a time
	for _, instruction := range instructions {
		currInstructionSet = strings.Split(instruction, " ")
		direction = toDirection(currInstructionSet[0])
		moves, _ = strconv.Atoi(currInstructionSet[1])

		for i := 0; i < moves; i++ {
			head.Move(direction)
			catch_up(*head, tail, &visited)
		}
	}
	return len(visited)
}

func main() {
	// read input
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print(fmt.Errorf("%w", err))
	}
	content_str := string(content)
	// 	content_str := `R 4
	// U 4
	// L 3
	// D 1
	// R 4
	// D 1
	// L 5
	// R 2`

	fmt.Println("Tail has visited: ", find_solution(content_str))
}
