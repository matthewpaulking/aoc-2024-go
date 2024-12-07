package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// file, err := os.Open("test.txt")
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	type Equation struct {
		total int
		nums  []int
		valid bool
	}

	scanner := bufio.NewScanner(file)
	test_cases := make([]Equation, 0)
	lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
		split := strings.Split(line, ":")
		test_case, err := strconv.Atoi(split[0])
		if err != nil {
			panic(err)
		}
		nums := make([]int, 0)
		for _, num := range strings.Split(split[1], " ") {
			if num == "" {
				continue
			}
			num, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			nums = append(nums, num)
		}
		test_cases = append(test_cases, Equation{total: test_case, nums: nums, valid: false})
	}

	// PART 1
	for i := range test_cases {
		equation := &test_cases[i]
		combos := operatorCombos(len(equation.nums) - 1)
		for _, val := range combos {
			total := 0
			for j := 0; j < len(val); j++ {
				if j == 0 {
					switch val[j] {
					case "+":
						total = equation.nums[j] + equation.nums[j+1]
					case "*":
						total = equation.nums[j] * equation.nums[j+1]
					}
				} else {
					switch val[j] {
					case "+":
						total = total + equation.nums[j+1]
					case "*":
						total = total * equation.nums[j+1]
					}
				}
			}
			if equation.total == total {
				equation.valid = true
				continue
			}
		}
	}
	answer1 := 0
	valids1 := 0
	for k := range test_cases {
		if test_cases[k].valid {
			answer1 = answer1 + test_cases[k].total
			valids1++
		}
	}
	fmt.Println("Answer 1:", answer1)
	fmt.Println(valids1, "out of", len(test_cases), "valid")

	// PART 2
	for i := range test_cases {
		equation := &test_cases[i]
		// equation.valid = false // reset to false
		combos := operatorCombos2(len(equation.nums) - 1)
		for _, val := range combos {
			total := 0
			for j := 0; j < len(val); j++ {
				if j == 0 {
					switch val[j] {
					case "+":
						total = equation.nums[j] + equation.nums[j+1]
					case "*":
						total = equation.nums[j] * equation.nums[j+1]
					case "||":
						total = concat(equation.nums[j], equation.nums[j+1])
					}
				} else {
					switch val[j] {
					case "+":
						total = total + equation.nums[j+1]
					case "*":
						total = total * equation.nums[j+1]
					case "||":
						total = concat(total, equation.nums[j+1])
					}
				}
			}
			if equation.total == total {
				equation.valid = true
				continue
			}
		}
	}

	answer2 := 0
	valids2 := 0
	for k := range test_cases {
		if test_cases[k].valid {
			answer2 = answer2 + test_cases[k].total
			valids2++
		} else {
			fmt.Println("Invalid:", test_cases[k])
		}
	}
	fmt.Println("Answer 2:", answer2)
	fmt.Println(valids2, "out of", len(test_cases), "valid")
}

func operatorCombos(spaces int) [][]string {
	total := 1 << spaces
	result := make([][]string, total)

	for i := 0; i < total; i++ {
		combo := make([]string, spaces)

		for j := 0; j < spaces; j++ {
			mask := 1 << j
			if (i & mask) == 0 {
				combo[j] = "+"
			} else {
				combo[j] = "*"
			}
		}
		result[i] = combo
	}
	return result
}

func operatorCombos2(spaces int) [][]string {
	operators := []string{"+", "*", "||"}
	total := 1 // Calculate 3^spaces
	for i := 0; i < spaces; i++ {
		total *= 3
	}

	result := make([][]string, total)

	for i := 0; i < total; i++ {
		combo := make([]string, spaces)
		num := i

		// Convert number to base-3 and map to operators
		for j := 0; j < spaces; j++ {
			combo[j] = operators[num%3]
			num /= 3
		}
		result[i] = combo
	}

	return result
}
func concat(a int, b int) int {
	x := strconv.Itoa(a)
	y := strconv.Itoa(b)
	combo := x + y
	num, err := strconv.Atoi(combo)
	if err != nil {
		panic(err)
	}
	return num
}
