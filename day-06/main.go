package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func checkBoundary(current_pos []int, grid [][]string, done *bool) bool {
	// fmt.Println("Current pos: ", curren_pos)
	// fmt.Println("Len X", len(grid[0]))
	// fmt.Println("Condition X: ", curren_pos[1] == 0 || curren_pos[0] > len(grid[0]))
	if current_pos[0] == 0 || current_pos[0] >= len(grid[0])-1 {
		fmt.Println("hit X end")
		*done = true
		return true
	}
	// fmt.Println("Len Y", len(grid[1]))
	// fmt.Println("Condition Y: ", curren_pos[1] == 0 || curren_pos[1] > len(grid[1]))
	if current_pos[1] == 0 || current_pos[1] >= len(grid[1])-1 {
		fmt.Println("hit Y end")
		*done = true
		return true
	}
	// fmt.Println("Current pos: ", curren_pos)
	return false
}

func goUp(current_pos []int, grid [][]string, direction *string, path [][]string) {
	// fmt.Println("UP")
	// fmt.Println("Current pos: ", current_pos)
	clone := slices.Clone(current_pos)
	clone[0] = clone[0] - 1
	target := grid[clone[0]][(clone[1])]
	// fmt.Println("---")
	// fmt.Println(grid[clone[0]][(clone[1])])
	// fmt.Println("---")
	p := fmt.Sprintf("Target (%d,%d): %s", clone[0], clone[1], target)
	fmt.Println(p)
	if target == "#" {
		fmt.Println("Hit obstacle")
		*direction = "right"

	} else {
		current_pos[0] = current_pos[0] - 1
		path[current_pos[0]][current_pos[1]] = "X"
	}
}
func goDown(current_pos []int, grid [][]string, direction *string, path [][]string) {
	fmt.Println("DOWN")
	clone := slices.Clone(current_pos)
	clone[0] = clone[0] + 1
	target := grid[clone[0]][clone[1]]
	p := fmt.Sprintf("Target (%d,%d): %s", clone[0], clone[1], target)
	fmt.Println(p)
	if target == "#" {
		fmt.Println("Hit obstacle")
		*direction = "left"
	} else {
		current_pos[0] = current_pos[0] + 1
		path[current_pos[0]][current_pos[1]] = "X"
	}
}
func goLeft(current_pos []int, grid [][]string, direction *string, path [][]string) {
	fmt.Println("LEFT")
	clone := slices.Clone(current_pos)
	target := grid[clone[0]][clone[1]-1]
	p := fmt.Sprintf("Target (%d,%d): %s", clone[0], clone[1], target)
	fmt.Println(p)
	if target == "#" {
		fmt.Println("Hit obstacle")
		*direction = "up"
	} else {
		current_pos[1] = current_pos[1] - 1
		path[current_pos[0]][current_pos[1]] = "X"
	}
}
func goRight(current_pos []int, grid [][]string, direction *string, path [][]string) {
	fmt.Println("RIGHT")
	clone := slices.Clone(current_pos)
	clone[1] = clone[1] + 1
	target := grid[clone[0]][clone[1]]
	p := fmt.Sprintf("Target (%d,%d): %s", clone[0], clone[1], target)
	fmt.Println(p)
	if target == "#" {
		fmt.Println("Hit obstacle")
		*direction = "down"
	} else {
		current_pos[1] = current_pos[1] + 1
		path[current_pos[0]][current_pos[1]] = "X"
	}
}

func main() {
	file, err := os.Open("input.txt")
	// file, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var grid [][]string
	var obstacles [][]int
	direction := "up"
	done := false
	current_pos := make([]int, 2)
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		split := strings.Split(line, "")
		for j := 0; j < len(split); j++ {
			if split[j] == "#" {
				obstacles = append(obstacles, []int{i, j})
			}
			if split[j] == "^" {
				current_pos[0] = i
				current_pos[1] = j
			}
		}
		grid = append(grid, strings.Split(line, ""))
	}

	path := slices.Clone(grid)
	path[current_pos[0]][current_pos[1]] = "X"
	fmt.Println("CURRENT POS: ", current_pos)

	for !done {
		fmt.Println("Direction: ", direction)
		if checkBoundary(current_pos, grid, &done) == true {
			break
		}
		if direction == "up" {
			fmt.Println("In Up")
			goUp(current_pos, grid, &direction, path)
			continue
		}
		if direction == "right" {
			fmt.Println("In Right")
			goRight(current_pos, grid, &direction, path)
			continue
		}
		if direction == "left" {
			fmt.Println("In Left")
			goLeft(current_pos, grid, &direction, path)
			continue
		}
		if direction == "down" {
			fmt.Println("In Down")
			goDown(current_pos, grid, &direction, path)
			continue
		}
		fmt.Println(current_pos)
	}

	// fmt.Println(done)
	// fmt.Println(direction)
	// fmt.Println(obstacles)
	// fmt.Println(current_pos)
	// fmt.Println(grid)
	fmt.Println(path)
	total_path := 0
	for i := 0; i < len(path); i++ {
		for j := 0; j < len(path[i]); j++ {
			if path[i][j] == "X" {
				total_path++
			}
		}
	}
	fmt.Println("TOTAL", total_path)
}
