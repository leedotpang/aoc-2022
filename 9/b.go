package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

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
	var dir Direction
	x_diff := int(math.Abs(float64(head.x - tail.x)))
	y_diff := int(math.Abs(float64(head.y - tail.y)))

	if x_diff < 2 && y_diff < 2 {
		return
	}
	if x_diff != 0 {
		dir = map[bool]Direction{true: RIGHT, false: LEFT}[head.x-tail.x > 0]
		tail.Move(dir)
	}
	if y_diff != 0 {
		dir = map[bool]Direction{true: UP, false: DOWN}[head.y-tail.y > 0]
		tail.Move(dir)
	}
}

func find_solution(content string, tail_count int) int {
	head := &Coordinate{x: 0, y: 0}
	tails := []*Coordinate{head}
	instructions := strings.Split(content, "\n")
	moves := 0
	currInstructionSet := []string{}
	var direction Direction
	visited := map[Coordinate]bool{*tails[0]: true}

	for i := 0; i < tail_count; i++ {
		tails = append(tails, &Coordinate{x: 0, y: 0})
	}

	last_tail_index := len(tails) - 1

	// Loop over instructions and move in required direction one move at a time
	for _, instruction := range instructions {
		currInstructionSet = strings.Split(instruction, " ")
		direction = toDirection(currInstructionSet[0])
		moves, _ = strconv.Atoi(currInstructionSet[1])

		for i := 0; i < moves; i++ {
			head.Move(direction)

			for ti, tail := range tails {
				if ti == 0 {
					continue
				}
				catch_up(*tails[ti-1], tail, &visited)
				if ti == last_tail_index {
					visited[*tail] = true
				}
			}
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

	fmt.Println("Tail has visited: ", find_solution(content_str, 9))
}
