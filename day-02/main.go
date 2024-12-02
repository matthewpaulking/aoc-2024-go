package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day-02.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// get current directory
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println(dir)

	scanner := bufio.NewScanner(file)
	var reports [][]int
	for scanner.Scan() {
		line := scanner.Text()

		ints := strings.Fields(line)
		report := make([]int, len(ints))
		for i, num := range ints {
			n, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			report[i] = n
		}
		reports = append(reports, report)
	}

	safe := 0
	for _, report := range reports {
		// fmt.Println(report)
		// fmt.Println(is_safe(report))
		if is_safe(report) {
			safe++
		}
	}
	fmt.Println(safe)

	safe = 0
	for _, report := range reports {
		if is_safe_with_dampener(report) {
			safe++
		}
	}
	fmt.Println(safe)

}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func direction(x int) int {
	switch {
	case x < 0:
		return -1
	case x > 0:
		return 1
	default:
		return 0
	}
}

func is_safe(report []int) bool {
	dir := direction(report[0] - report[1])
	for i := 0; i < len(report)-1; i++ {
		current_dir := direction(report[i] - report[i+1])
		if current_dir != dir {
			return false
		}
		diff := abs(report[i] - report[i+1])
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func is_safe_with_dampener(report []int) bool {
	if is_safe(report) {
		return true
	}
	for i := range report {
		clone := slices.Clone(report)
		check := slices.Delete(clone, i, i+1)
		if is_safe(check) {
			return true
		}
	}
	return false
}
