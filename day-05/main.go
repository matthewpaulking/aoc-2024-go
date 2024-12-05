package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse_integers(s []string) ([]int, error) {
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

func main() {
	// file, err := os.Open("test.txt")
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rules := make([][]int, 0)
	rule_map := make(map[int][]int)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		parts := strings.Split(line, "|")
		num1, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		num2, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		rule_map[num1] = append(rule_map[num1], num2)
		rules = append(rules, []int{num1, num2})
	}

	changes := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		parts := strings.Split(line, ",")
		nums, err := parse_integers(parts)
		if err != nil {
			panic(err)
		}

		changes = append(changes, nums)
	}

	results := make([]bool, 0)
	for _, change := range changes {
		good := true

		for i, page := range change {
			if !good {
				break
			}
			rules := rule_map[page]
			// check the "afters"
			for j := i + 1; j < len(change); j++ {
				for _, rule := range rules {
					if rule == change[j] {
						good = true
						break
					}
				}
			}
			//check the "befores"
			if i > 0 {
				for j := i - 1; j >= 0; j-- {
					for _, rule := range rules {
						if rule == change[j] {
							good = false
							break
						}
					}
				}
			}
		}
		results = append(results, good)
	}

	answer1 := 0
	for i, result := range results {
		if result == true {
			// find middle number
			length := len(changes[i])
			if length%2 == 0 {
				panic("length is even")
			}
			middle := length / 2
			answer1 += changes[i][middle]
		}
	}
	fmt.Println(results)

	// Part 2
	for i, change := range changes {
		if results[i] == true {
			continue
		}
		fmt.Println("***")
		fmt.Println(change)
		fmt.Println("")
		fmt.Println("")
		for p, page := range change {
			rules := rule_map[page]
			if p > 0 {
				for k := p - 1; k >= 0; k-- {
					for _, rule := range rules {
						if rule == change[k] {
							if p > 0 {
								diff := p - k
								change[k+1], change[p-diff] = change[p-diff], change[k+1]
							}
					}
				}
			}
		}
		fmt.Println("new change: ", change)
	}

	answer2 := 0
	for i, change := range changes {
		if results[i] == true {
			continue
		}
		length := len(change)
		middle := length / 2
		answer2 += change[middle]
	}

	fmt.Println("Answer 1: ", answer1)
	fmt.Println("Answer 2: ", answer2)
}
