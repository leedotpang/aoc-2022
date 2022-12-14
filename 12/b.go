package main

import (
	"fmt"
	"lee/aoc/helpers"
	"os"
	"strings"
)

type Coordinates struct {
	y int
	x int
}

type SearchQueue = helpers.Stack[Coordinates]
type Connections = helpers.Stack[Coordinates]

func find_many_coordinates(s rune, rows []string) []Coordinates {
	coordinates := SearchQueue{}

	for vi, row := range rows {
		for hi, col := range row {
			if col == s {
				coordinates.Push(Coordinates{x: hi, y: vi})
			}
		}
	}

	return coordinates
}
func find_coordinates(s rune, rows []string) Coordinates {
	for vi, row := range rows {
		for hi, col := range row {
			if col == s {
				return Coordinates{x: hi, y: vi}
			}
		}
	}

	return Coordinates{}
}

func calculate_shortest_path(s Coordinates, e Coordinates, rows []string) int {
	curr_coordinate := s
	search_queue := SearchQueue{s}
	connections := map[Coordinates]*SearchQueue{}
	distances := map[Coordinates]int{}
	visited := map[Coordinates]bool{}

	for vi, row := range rows {
		for hi, col := range row {
			curr_coordinate = Coordinates{x: hi, y: vi}
			connections[curr_coordinate] = &Connections{}
			distances[curr_coordinate] = 99999

			if curr_coordinate == s {
				distances[curr_coordinate] = 0
			}
			if curr_coordinate == e {
				col = 'z'
			}

			// Right
			if hi+1 < len(row) && col >= rune(row[hi+1]-1) {
				connections[curr_coordinate].Push(Coordinates{x: hi + 1, y: vi})
			}
			// Left
			if hi-1 >= 0 && col >= rune(row[hi-1]-1) {
				connections[curr_coordinate].Push(Coordinates{x: hi - 1, y: vi})
			}
			// Up
			if vi+1 < len(rows) && col >= rune(rows[vi+1][hi]-1) {
				connections[curr_coordinate].Push(Coordinates{x: hi, y: vi + 1})
			}
			// Down
			if vi-1 >= 0 && col >= rune(rows[vi-1][hi]-1) {
				connections[curr_coordinate].Push(Coordinates{x: hi, y: vi - 1})
			}
		}
	}

	for q_len := len(search_queue); q_len > 0; q_len = len(search_queue) {
		next := search_queue.Shift()

		for _, node := range *connections[next] {
			if _, seen := visited[node]; !seen {
				search_queue.Push(node)
				visited[node] = true
			}

			if distances[node] > distances[next]+1 {
				distances[node] = distances[next] + 1
			}
		}
	}

	return distances[e]
}

func find_solution(content string) int {
	rows := strings.Split(content, "\n")
	start_coordinates := find_many_coordinates('a', rows)
	end_coordinates := find_coordinates('E', rows)
	shortest_path := 99999
	curr_val := 0

	for _, start := range start_coordinates {
		curr_val = calculate_shortest_path(start, end_coordinates, rows)
		if curr_val < shortest_path {
			shortest_path = curr_val
		}
	}

	return shortest_path
}

func main() {
	// read input
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print(fmt.Errorf("%w", err))
	}
	content_str := string(content)

	fmt.Println("Fewest steps: ", find_solution(content_str))
}
