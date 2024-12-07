package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Point struct {
	x, y int
}

type Direction int

const (
	Up    Direction = 0
	Right Direction = 1
	Down  Direction = 2
	Left  Direction = 3
)

func is_out_of_bounds(pos Point, grid [][]string) bool {
	return pos.x <= 0 || pos.x >= len(grid)-1 ||
		pos.y <= 0 || pos.y >= len(grid[0])-1
}

func move(pos Point, grid [][]string, dir *Direction, path [][]string) Point {
	new_pos := pos

	switch *dir {
	case Up:
		new_pos.x--
	case Right:
		new_pos.y++
	case Down:
		new_pos.x++
	case Left:
		new_pos.y--
	}

	// Check if hitting obstacle
	if grid[new_pos.x][new_pos.y] == "#" {
		// Rotate direction clockwise
		*dir = func() Direction {
			switch *dir {
			case Up:
				return Right
			case Right:
				return Down
			case Down:
				return Left
			case Left:
				return Up
			default:
				return Up
			}
		}()
		return pos
	}

	path[new_pos.x][new_pos.y] = "X"
	return new_pos
}

func move_with_loop_detect(pos Point, grid [][]string, dir *Direction, path [][]string, loop *[]Point, first_obstacle_found *bool) Point {
	new_pos := pos

	switch *dir {
	case Up:
		new_pos.x--
	case Right:
		new_pos.y++
	case Down:
		new_pos.x++
	case Left:
		new_pos.y--
	}

	// Check if hitting obstacle
	if grid[new_pos.x][new_pos.y] == "#" {
		// fmt.Println("OBSTACLE FOUND: ", new_pos)
		*first_obstacle_found = true
		*loop = append(*loop, new_pos)
		// Rotate direction clockwise
		*dir = func() Direction {
			switch *dir {
			case Up:
				return Right
			case Right:
				return Down
			case Down:
				return Left
			case Left:
				return Up
			default:
				return Up
			}
		}()
		return pos
	}

	path[new_pos.x][new_pos.y] = "X"
	return new_pos
}

func count_path(path [][]string) int {
	total := 0
	for i := range path {
		for j := range path[i] {
			if path[i][j] == "X" {
				total++
			}
		}
	}
	return total
}

func read_grid(filename string) ([][]string, Point) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var grid [][]string
	var start_pos Point

	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		row := strings.Split(line, "")

		// Find starting position
		for j, char := range row {
			if char == "^" {
				start_pos = Point{i, j}
			}
		}

		grid = append(grid, row)
	}

	return grid, start_pos
}

func main() {
	// grid, pos := read_grid("test.txt")
	grid, pos := read_grid("input.txt")
	original_grid := deep_clone_grid(grid)
	starting_pos := pos
	path := slices.Clone(grid)
	dir := Up

	// Mark starting position
	path[pos.x][pos.y] = "X"

	// Main movement loop
	for !is_out_of_bounds(pos, grid) {
		pos = move(pos, grid, &dir, path)
	}

	total := count_path(path)
	fmt.Println("Answer 1:", total)

	obstruction_count := 0
	for i := range path {
		for j := range path[i] {
			// Skip starting position
			if i == starting_pos.x && j == starting_pos.y {
				continue
			}
			// Skip grid positions where guard will not go
			if path[i][j] != "X" {
				continue
			}
			// add obstacle to copy of grid
			grid_copy := deep_clone_grid(grid)
			current_dir := Up
			grid_copy[i][j] = "#"
			grid_copy[starting_pos.x][starting_pos.y] = "^"
			pos = starting_pos
			loop := make([]Point, 0)
			first_obstacle_found := false
			for !is_out_of_bounds(pos, original_grid) {
				if first_obstacle_found && detect_loop(loop) {
					obstruction_count++
					break
				}
				pos = move_with_loop_detect(pos, grid_copy, &current_dir, path, &loop, &first_obstacle_found)
			}
		}
	}

	fmt.Println("Answer 2:", obstruction_count)
}
func deep_clone_grid(grid [][]string) [][]string {
	new_grid := make([][]string, len(grid))
	for i := range grid {
		new_grid[i] = make([]string, len(grid[i]))
		copy(new_grid[i], grid[i])
	}
	return new_grid
}

func detect_loop(points []Point) bool {
	if len(points) > 150 { // set this arbitrarily high
		return true
	}
	return false
}
