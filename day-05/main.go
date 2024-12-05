package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Helper function to convert string slice to integers
func parseIntegers(s []string) ([]int, error) {
	result := make([]int, 0, len(s))
	for _, str := range s {
		num, err := strconv.Atoi(strings.TrimSpace(str))
		if err != nil {
			return nil, err
		}
		result = append(result, num)
	}
	return result, nil
}

// Parse the rules section of the input
func parseRules(scanner *bufio.Scanner) (map[int][]int, error) {
	ruleMap := make(map[int][]int)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		parts := strings.Split(line, "|")
		num1, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, err
		}
		num2, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}
		ruleMap[num1] = append(ruleMap[num1], num2)
	}
	return ruleMap, nil
}

// Parse the changes section of the input
func parseChanges(scanner *bufio.Scanner) ([][]int, error) {
	changes := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		nums, err := parseIntegers(strings.Split(line, ","))
		if err != nil {
			return nil, err
		}
		changes = append(changes, nums)
	}
	return changes, nil
}

// Check if a change sequence is valid according to rules
func isValidChange(change []int, ruleMap map[int][]int) bool {
	for i, page := range change {
		rules := ruleMap[page]
		// Check if any number that should come after appears before
		for j := 0; j < i; j++ {
			if contains(rules, change[j]) {
				return false
			}
		}
	}
	return true
}

// Helper function to check if a slice contains a value
func contains(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

// Get middle value from a slice
func getMiddleValue(slice []int) int {
	return slice[len(slice)/2]
}

// Fix the order of a change sequence
func fixChangeOrder(change []int, ruleMap map[int][]int) []int {
	result := make([]int, len(change))
	copy(result, change)

	for i := 1; i < len(result); i++ {
		for j := i; j > 0; j-- {
			curr := result[j]
			prev := result[j-1]
			if contains(ruleMap[curr], prev) {
				// Swap if current number should come before previous
				result[j], result[j-1] = result[j-1], result[j]
			}
		}
	}
	return result
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	ruleMap, err := parseRules(scanner)
	if err != nil {
		panic(err)
	}

	changes, err := parseChanges(scanner)
	if err != nil {
		panic(err)
	}

	// Part 1
	answer1 := 0
	validChanges := make([]bool, len(changes))
	for i, change := range changes {
		validChanges[i] = isValidChange(change, ruleMap)
		if validChanges[i] {
			answer1 += getMiddleValue(change)
		}
	}

	// Part 2
	answer2 := 0
	for i, change := range changes {
		if !validChanges[i] {
			fixedChange := fixChangeOrder(change, ruleMap)
			answer2 += getMiddleValue(fixedChange)
		}
	}

	fmt.Println("Answer 1:", answer1)
	fmt.Println("Answer 2:", answer2)
}
