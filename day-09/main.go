package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	// "slices"
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
func blocks_to_string(b []Block) []string {
	output := make([]string, 0)
	for _, v := range b {
		if v.id == -1 {
			output = append(output, fill(".", v.length)...)
		} else {
			output = append(output, fill(convert_to_string(v.id), v.length)...)
		}
	}
	return output
}

func main() {
	file, err := os.Open("test.txt")
	// file, err := os.Open("input.txt")
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
	// memory_string_2 := slices.Clone(memory_string)

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
	// fmt.Println(blocks)
	// fmt.Println(blocks_to_string(blocks))

	// SWAPPING BLOCKS DOESN'T WORK

	tick := 0
	for i, v := range blocks {
		if v.length == 0 || v.id >= 0 {
			continue
		}
		// HERE
		for j := id_count; j >= 0; j-- {
			swap := len(blocks) - 1
			for blocks[swap].id == -1 {
				swap--
			}
			for j := swap; j >= 0; j-- {
				// fmt.Println("SWAP BLOCK", blocks[swap])
				fmt.Println("BLOCK", v, "SWAP", blocks[swap])
				diff := v.length - blocks[swap].length
				if diff >= 0 {
					// split the block
					staying := Block{id: -1, length: diff}
					swapping := Block{id: -1, length: 3 - diff}
					blocks[i], blocks[swap] = blocks[swap], swapping
					swap--
					// insert staying block after i
					blocks = append(blocks[:i+1], append([]Block{staying}, blocks[i+1:]...)...)
					i++
					tick++
				} else {
					// move on to next block
					// fmt.Println("DIFF IS GREATER")
					swap--
				}
			}
		}
		fmt.Println(blocks_to_string(blocks))
	}
	fmt.Println(blocks)
	fmt.Println(blocks_to_string(blocks))

	// STRING APPROACH TO PART2
	// swap := len(memory_string_2) - 1
	// for i, v := range memory_string_2 {
	// 	if swap == i {
	// 		break
	// 	}
	// 	if v == "." {
	// 		for memory_string_2[swap] == "." {
	// 			swap--
	// 		}
	// 		memory_string_2[i], memory_string_2[swap] = memory_string_2[swap], memory_string_2[i]
	// 		swap--
	// 	}
	// }

	// total2 := 0
	// for i, v := range blocks {
	// 	if v.id == -1 {
	// 		break
	// 	}
	// 	total1 = total1 + (i * convert_to_int(v))
	// }
}
