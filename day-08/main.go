package main

import (
	"bufio"
	"fmt"
	"os"
)

type Distance struct {
	dx int
	dy int
}
type Position struct {
	x int
	y int
}

func distance(a, b Position) Distance {
	return Distance{dx: a.x - b.x, dy: a.y - b.y}
}

func out_of_bounds(grid_length int, p Position) bool {
	if p.x < 0 || p.x >= grid_length {
		return true
	}
	if p.y < 0 || p.y >= grid_length {
		return true
	}
	return false
}

func main() {
	file, err := os.Open("input.txt")
	// file, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	frequencies := make(map[byte][]Position)
	antinodes := make(map[Position]int)
	grid_length := 0
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		if i == 0 {
			grid_length += len(line)
		}
		for j := 0; j < len(line); j++ {
			if line[j] != '.' {
				frequencies[line[j]] = append(frequencies[line[j]], Position{x: j, y: i})
			}
		}
	}
	// PART 1
	for _, v := range frequencies {
		for i := 0; i < len(v)-1; i++ {
			for j := i + 1; j < len(v); j++ {
				pos1 := v[i]
				pos2 := v[j]
				dist := distance(pos1, pos2)
				antinode1 := Position{pos1.x + dist.dx, pos1.y + dist.dy}
				if !out_of_bounds(grid_length, antinode1) {
					antinodes[antinode1] = antinodes[antinode1] + 1
				}
				antinode2 := Position{pos2.x - dist.dx, pos2.y - dist.dy}
				if !out_of_bounds(grid_length, antinode2) {
					antinodes[antinode2] = antinodes[antinode2] + 1
				}
			}
		}
	}
	fmt.Println("ANSWER 1:", len(antinodes))

	// PART 2 - Antinodes keep going until out of bounds
	for _, v := range frequencies {
		for i := 0; i < len(v)-1; i++ {
			for j := i + 1; j < len(v); j++ {
				pos1 := v[i]
				pos2 := v[j]
				// Always write in the first one
				antinodes[pos1] = antinodes[pos1] + 1
				dist := distance(pos1, pos2)
				cur_antinode1 := Position{pos1.x + dist.dx, pos1.y + dist.dy}
				for !out_of_bounds(grid_length, cur_antinode1) {
					antinodes[cur_antinode1] = antinodes[cur_antinode1] + 1
					cur_antinode1 = Position{cur_antinode1.x + dist.dx, cur_antinode1.y + dist.dy}
				}
				cur_antinode2 := Position{pos1.x - dist.dx, pos1.y - dist.dy}
				for !out_of_bounds(grid_length, cur_antinode2) {
					antinodes[cur_antinode2] = antinodes[cur_antinode2] + 1
					cur_antinode2 = Position{cur_antinode2.x - dist.dx, cur_antinode2.y - dist.dy}
				}
			}
		}
	}
	fmt.Println("ANSWER 2:", len(antinodes))
}
