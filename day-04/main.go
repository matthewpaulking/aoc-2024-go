package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// file, err := os.Open("day-04-warmup.txt")
	file, err := os.Open("day-04.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var puzzle []string
	for scanner.Scan() {
		puzzle = append(puzzle, scanner.Text())
	}

	answer1 := 0
	for i := 0; i < len(puzzle); i++ {
		for j, s := range puzzle[i] {
			if s == 'X' {
				answer1 += search(puzzle, i, j)
			}
		}
	}

	answer2 := 0
	for i := 0; i < len(puzzle); i++ {
		for j, s := range puzzle[i] {
			if s == 'A' {
				answer2 += xmas(puzzle, i, j)
			}
		}
	}

	fmt.Println("Answer 1: ", answer1)
	fmt.Println("Answer 2: ", answer2)
}

func search(s []string, i, j int) int {
	total := 0
	l := len(s)

	// HORIZONTAL
	if j <= (l - 1 - 3) {
		if s[i][j] == byte('X') && s[i][j+1] == byte('M') && s[i][j+2] == byte('A') && s[i][j+3] == byte('S') {
			total++
		}
	}
	// HORIZONTAL REVERSE
	if j >= 3 {
		if s[i][j] == byte('X') && s[i][j-1] == byte('M') && s[i][j-2] == byte('A') && s[i][j-3] == byte('S') {
			total++
		}
	}
	// VERTICAL
	if i <= (l - 1 - 3) {
		if s[i][j] == byte('X') && s[i+1][j] == byte('M') && s[i+2][j] == byte('A') && s[i+3][j] == byte('S') {
			total++
		}
	}
	// VERTICAL REVERSE
	if i >= 3 {
		if s[i][j] == byte('X') && s[i-1][j] == byte('M') && s[i-2][j] == byte('A') && s[i-3][j] == byte('S') {
			total++
		}
	}
	// DIAG RIGHT DOWN
	if i <= (l-1-3) && j <= (l-1-3) {
		if s[i][j] == byte('X') && s[i+1][j+1] == byte('M') && s[i+2][j+2] == byte('A') && s[i+3][j+3] == byte('S') {
			total++
		}
	}
	// DIAG RIGHT UP
	if i >= 3 && j <= (l-1-3) {
		if s[i][j] == byte('X') && s[i-1][j+1] == byte('M') && s[i-2][j+2] == byte('A') && s[i-3][j+3] == byte('S') {
			total++
		}
	}
	// DIAG LEFT DOWN
	if i <= (l-1-3) && j >= 3 {
		if s[i][j] == byte('X') && s[i+1][j-1] == byte('M') && s[i+2][j-2] == byte('A') && s[i+3][j-3] == byte('S') {
			total++
		}
	}
	// DIAG LEFT UP
	if i >= 3 && j >= 3 {
		if s[i][j] == byte('X') && s[i-1][j-1] == byte('M') && s[i-2][j-2] == byte('A') && s[i-3][j-3] == byte('S') {
			total++
		}
	}

	return total
}

func xmas(s []string, i, j int) int {
	total := 0
	l := len(s)

	if (i < 1 || i > l-2) || (j < 1 || j > l-2) {
		return 0
	}

	// S.S
	// .A.
	// M.M
	// top-left                    top-right                   bottom-left                 bottom-right
	if s[i-1][j-1] == byte('S') && s[i-1][j+1] == byte('S') && s[i+1][j-1] == byte('M') && s[i+1][j+1] == byte('M') {
		total++
	}

	// M.S
	// .A.
	// M.S
	// top-left                    top-right                   bottom-left                 bottom-right
	if s[i-1][j-1] == byte('M') && s[i-1][j+1] == byte('S') && s[i+1][j-1] == byte('M') && s[i+1][j+1] == byte('S') {
		total++
	}

	// M.M
	// .A.
	// S.S
	// top-left                    top-right                   bottom-left                 bottom-right
	if s[i-1][j-1] == byte('M') && s[i-1][j+1] == byte('M') && s[i+1][j-1] == byte('S') && s[i+1][j+1] == byte('S') {
		total++
	}

	// S.M
	// .A.
	// S.M
	// top-left                    top-right                   bottom-left                 bottom-right
	if s[i-1][j-1] == byte('S') && s[i-1][j+1] == byte('M') && s[i+1][j-1] == byte('S') && s[i+1][j+1] == byte('M') {
		total++
	}
	return total
}
