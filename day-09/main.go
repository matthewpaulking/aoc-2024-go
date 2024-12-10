package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
)

type Block struct {
	id     int
	length int
}

func convert_to_int(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return num
}

func convert_to_string(i int) string {
	num := strconv.Itoa(i)
	return num
}

func fill(s string, n int) []string {
	result := make([]string, n)
	for i := range result {
		result[i] = s
	}
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
	blocks := make([]Block, 0)
	file_list := list.New()
	memory_list := list.New()
	memory_list_with_spaces := list.New()
	id_count := 0
	memory_string := make([]string, 0)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		for j := 0; j < len(line); j++ {
			num := convert_to_int(string(line[j]))
			if j != 0 && j%2 != 0 {
				// fmt.Println("FREE SPACE LENGTH:", num)
				block := Block{id: -1, length: num}
				blocks = append(blocks, block)
				file_list.PushBack(num)
				memory_string = append(memory_string, fill(".", num)...)
				for k := 0; k < num; k++ {
					memory_list_with_spaces.PushBack(".")
				}
			} else {
				// fmt.Println("FILE BLOCK LENGTH:", num)
				block := Block{id: id_count, length: num}
				blocks = append(blocks, block)
				file_list.PushBack(num)
				for k := 0; k < num; k++ {
					memory_list.PushFront(strconv.Itoa(id_count))
					memory_list_with_spaces.PushBack(strconv.Itoa(id_count))
				}
				memory_string = append(memory_string, fill(strconv.Itoa(id_count), num)...)
				id_count++
			}
		}
	}
	// fmt.Println(memory_string)
	swap_index := len(memory_string) - 1
	for i, v := range memory_string {
		if swap_index == i {
			break
		}
		if v == "." {
			for memory_string[swap_index] == "." {
				swap_index--
			}
			memory_string[i], memory_string[swap_index] = memory_string[swap_index], memory_string[i]
			swap_index--
		}
		// fmt.Println(memory_string)

	}
	// fmt.Println(memory_string)
	total1 := 0
	for i, v := range memory_string {
		if v == "." {
			break
		}
		total1 = total1 + (i * convert_to_int(v))
	}
	fmt.Println("Answer 1:", total1)
}
