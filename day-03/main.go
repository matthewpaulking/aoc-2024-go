package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func extractNumberPair(s string) (int, int) {
	re := regexp.MustCompile(`(\d{1,3})`)
	numbers := re.FindAllString(s, -1)
	number_1, err := strconv.Atoi(numbers[0])
	if err != nil {
		panic(err)
	}
	number_2, err := strconv.Atoi(numbers[1])
	if err != nil {
		panic(err)
	}
	return number_1, number_2
}

func main() {
	input, err := os.ReadFile("day-03.txt")
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(input))
	re, err := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)
	if err != nil {
		panic(err)
	}

	out := re.FindAllString(string(input), -1)
	answer1 := 0
	for _, s := range out {
		number_1, number_2 := extractNumberPair(s)
		answer1 += number_1 * number_2
	}
	fmt.Printf("Answer 1: %d\n", answer1)

	re2, err := regexp.Compile(`(do\(\))|(mul\(\d{1,3},\d{1,3}\))|(don't\(\))`)
	out2 := re2.FindAllString(string(input), -1)
	accept := true
	accepted_operations := []string{}
	for i, s := range out2 {
		if s == "do()" {
			accept = true
			continue
		}
		if s == "don't()" {
			accept = false
			continue
		}
		if accept {
			accepted_operations = append(accepted_operations, out2[i])
		}
	}

	answer2 := 0
	for _, s := range accepted_operations {
		number_1, number_2 := extractNumberPair(s)
		answer2 += number_1 * number_2
	}
	fmt.Printf("Answer 2: %d\n", answer2)
}
