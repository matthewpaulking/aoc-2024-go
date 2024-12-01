package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
	}
	scanner := bufio.NewScanner(file)
	var col1, col2 []int
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		// fields[0] is first column, fields[1] is second column
		num1, err := strconv.Atoi(fields[0])
		if err != nil {
		}
		num2, err := strconv.Atoi(fields[1])
		if err != nil {
		}
		col1 = append(col1, num1)
		col2 = append(col2, num2)
	}
	sort.Ints(col1)
	sort.Ints(col2)

	diffs := 0
	for i := 0; i < len(col1); i++ {
		diff := col1[i] - col2[i]
		if diff < 0 {
			diff = -diff
		}
		diffs = diffs + diff
	}
	fmt.Println(diffs)

	counts := make(map[int]int)
	for _, num := range col2 {
		counts[num]++
	}
	results := make([]int, len(col1))
	for i, num := range col1 {
		results[i] = counts[num]
	}
	// fmt.Println(results)

	answer2 := 0
	for i := 0; i < len(col1); i++ {
		mult := col1[i] * results[i]
		answer2 = answer2 + mult
	}
	fmt.Println(answer2)
}
