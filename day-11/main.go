package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func split(i int, memo *map[int][]int) []int {
	// Check if we've seen this stone before
	if _, ok := (*memo)[i]; ok {
		// fmt.Println("Found in memo:", i)
		return (*memo)[i]
	}

	if i == 0 {
		result := []int{1}
		(*memo)[i] = result
		return result
	}

	digits := strings.Split(strconv.Itoa(i), "")
	if len(digits)%2 == 0 {
		l := len(digits)
		left := strings.Join(digits[0:l/2], "")
		left_int, _ := strconv.Atoi(left)
		right := strings.Join(digits[l/2:], "")
		right_int, _ := strconv.Atoi(right)
		result := []int{left_int, right_int}
		return result
	}
	result := []int{i * 2024}

	// Store the result in memo before returning
	(*memo)[i] = result
	return result
}

func main() {
	// file, err := os.Open("test.txt")
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	stones := make([]int, 0)
	memo := make(map[int][]int)

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		for _, part := range line {
			num, _ := strconv.Atoi(part)
			stones = append(stones, num)
		}
	}

	// Answer 1 - not using a map.
	current_stones := slices.Clone(stones)
	iters := 25
	// answer1 := make([]int, 0)
	for i := 0; i < iters; i++ {
		new_stones := make([]int, 0)
		for _, stone := range current_stones {
			// fmt.Println("Splitting:", stone)
			s := split(stone, &memo)
			new_stones = append(new_stones, s...)
			// fmt.Println("Result:", s)
		}
		current_stones = new_stones
	}
	fmt.Println("Answer1:", len(current_stones))

	// Answer 2 - gotta use a map. I assume it's rebuilding a giant array that's costly.
	final_stones := make(map[int]int, 0)
	mapped_stones := make(map[int]int, 0)
	for _, stone := range stones {
		mapped_stones[stone] = 1
	}
	iters = 75
	for i := 0; i < iters; i++ {
		new_stones := make(map[int]int, 0)
		for k, v := range mapped_stones {
			s := split(k, &memo)
			for _, s2 := range s {
				new_stones[s2] = new_stones[s2] + v
			}
		}
		mapped_stones = new_stones
		if i == iters-1 {
			for k, v := range new_stones {
				final_stones[k] = v
			}
		}
	}
	answer2 := 0
	if len(final_stones) > 0 {
		for _, v := range final_stones {
			answer2 += v
		}
	}
	fmt.Println("Answer2:", answer2)
	// fmt.Println(final_stones)
}
