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

type SearchQueue = helpers.Queue[Coordinates]
type Connections = helpers.Stack[Coordinates]

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
	connections := map[Coordinates]*Connections{}
	distances := map[Coordinates]int{}
	visited := map[Coordinates]bool{}

	// Create edges/connections
	for vi, row := range rows {
		for hi, col := range row {
			curr_coordinate = Coordinates{x: hi, y: vi}
			connections[curr_coordinate] = &Connections{}
			distances[curr_coordinate] = 99999

			if curr_coordinate == s {
				distances[curr_coordinate] = 0
				col = 'a'
			}
			if curr_coordinate == e {
				col = 'z'
			}

			// Look in four directions and if path exists create it

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

	// Loop through paths and calculate shortest path to each node given each other possible node
	// yes...dijkstra's algo
	for search_queue.Present() {
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
	// get coordinates of S and E
	start_coordinates := find_coordinates('S', rows)
	end_coordinates := find_coordinates('E', rows)

	// determine shortest path between them
	shortest_path := calculate_shortest_path(start_coordinates, end_coordinates, rows)

	return shortest_path
}

/*
It may be worth revisiting but this is the first day where my solution worked for both parts, but not for the example input... there is an error somewhere in my logic --> but I guess I'll revisit this in the future and hopefully laugh it off. For now, I'll just let this keep me up at night.
*/
func main() {
	// read input
	// start := time.Now()
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print(fmt.Errorf("%w", err))
	}
	content_str := string(content)

	fmt.Println("Fewest steps: ", find_solution(content_str))
	// fmt.Println(time.Since(start).Seconds())
}
