// package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"
)

func count_xmas(s string) int {
	re := regexp.MustCompile(`XMAS`)
	matches := re.FindAllString(s, -1)
	return len(matches)
}

func reverse(s string) string {
	reverse := strings.Split(s, "")
	slices.Reverse(reverse)
	return strings.Join(reverse, "")
}
func makeRange(start, end int) []int {
	diff := end - start
	if diff < 0 {
		diff = -diff
	}
	nums := make([]int, 0)
	if start == end {
		nums = append(nums, start)
		return nums
	}
	if start > end {
		for i := start; i >= end; i-- {
			nums = append(nums, i)
		}
	} else {
		for i := start; i <= end; i++ {
			nums = append(nums, i)
		}
	}
	return nums
}
func get_diag(s [][]string, start int, end int) []string {
	diff := start - end
	if diff < 0 {
		diff = -diff
	}
	// fmt.Println("diff: ", diff)
	out_string := make([]string, diff+1)
	// fmt.Println(out_string)
	if start > end {
		for i := start; i >= end; i-- {
			// fmt.Println("i: ", i)
			i_norm := start - i
			forward := makeRange(start, i)
			backward := makeRange(i, start)
			// fmt.Println(forward)
			// fmt.Println(backward)
			for j := range forward {
				// fmt.Println("J == ", j)
				// fmt.Println("Forward", forward[j])
				// fmt.Println("Backward", backward[j])
				out_string[i_norm] = out_string[i_norm] + s[forward[j]][backward[j]]
			}
			// fmt.Println("------------------")
			// fmt.Println("")
		}
	} else {
		for i := start; i < end; i++ {
			forward := makeRange(start, i)
			backward := makeRange(i, start)
			for j := range forward {
				// fmt.Println("J == ", j)
				// fmt.Println("Forward", forward[j])
				// fmt.Println("Backward", backward[j])
				// out_string = append(out_string, s[forward[j]][backward[j]])
				out_string[i] = out_string[i] + s[forward[j]][backward[j]]
			}
			// fmt.Println("------------------")
			// fmt.Println("")
		}
	}
	// final := strings.Join(out_string, "")
	return out_string
}

func main() {
	file, err := os.Open("day-04-warmup.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	horizontal := ""
	raw := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		raw = append(raw, line)
		horizontal += line
		// split := strings.Split(line, "")
		// reversed := slices.Clone(split)
		// slices.Reverse(reversed)
		// fmt.Println(line)
		// fmt.Printf("%T\n", line)
		// fmt.Println(split)
		// fmt.Printf("%T\n", split)
		// fmt.Println(reversed)
		// fmt.Printf("%T\n", reversed)
		// reversed_string := strings.Join(reversed, "")
		// fmt.Println(reversed_string)
		// fmt.Printf("%T\n", reversed_string)

		// revline := line
		// occ := len(re.FindAllString(line, -1))
		// fmt.Printf("Regular count: %s\n", occ)
		// fmt.Printf("Reverse count: %s\n", len(reverse))
	}

	// VERTICAL
	line_length := len(raw[0])
	vertical := make([]string, line_length)
	for i := 0; i < line_length; i++ {
		split := strings.Split(raw[i], "")
		for j := range split {
			vertical[j] = vertical[j] + split[j]
		}
	}

	// vertical = strings.Join(vertical, "")
	fmt.Println("--- VERTICAL ---")
	vert_combined := strings.Join(vertical, "")
	reversed_vertical := reverse(vert_combined)
	// fmt.Println(vert_combined)
	// fmt.Println(reversed_vertical)
	vertical_1 := fmt.Sprintf("Norm: %d", count_xmas(vert_combined))
	vertical_2 := fmt.Sprintf("Rev : %d", count_xmas(reversed_vertical))
	fmt.Println(vertical_1)
	fmt.Println(vertical_2)

	reversed := reverse(horizontal)
	horizontal_1 := fmt.Sprintf("Norm: %d", count_xmas(horizontal))
	horizontal_2 := fmt.Sprintf("Rev : %d", count_xmas(reversed))
	fmt.Println("--- HORIZONTAL ---")
	fmt.Println(horizontal_1)
	fmt.Println(horizontal_2)

	raw_split := make([][]string, 0)
	raw_split_rev := make([][]string, 0)

	for _, s := range raw {
		sp := strings.Split(s, "")
		rev := slices.Clone(sp)
		slices.Reverse(rev)
		raw_split = append(raw_split, sp)
		raw_split_rev = append(raw_split_rev, rev)
	}
	// fmt.Println(raw_split)
	// fmt.Println(raw_split_rev)
	// for _, s := range raw_split {
	// 	fmt.Println(s)
	// }

	// vertical := make([]string, len(raw))
	// for i, s := range raw {
	// 	split := strings.Split(s, "")
	// 	fmt.Println(split)
	// 	vertical[i] = vertical[i] + split[i]
	// }
	// fmt.Println(vertical)

	// fmt.Println(makeRange(0, 9))
	// fmt.Println(makeRange(9, 0))
	// fmt.Println(makeRange(9, 9))
	// fmt.Println(makeRange(8, 9))
	// fmt.Println(makeRange(1, 2))

	// right diags
	right_diags := append(get_diag(raw_split, 0, 9), get_diag(raw_split, 9, 1)...)
	// fmt.Println(get_diag(raw_split, 0, 9))
	// fmt.Println(get_diag(raw_split, 9, 1))
	// fmt.Println(right_diags)
	right_diags_reversed := make([]string, len(right_diags))
	for i, s := range right_diags {
		split := strings.Split(s, "")
		slices.Reverse(split)
		right_diags_reversed[i] = strings.Join(split, "")
	}
	// now find XMAS in every item in the slices
	right_diag_sum := 0
	for _, s := range right_diags {
		right_diag_sum += count_xmas(s)
	}
	for _, s := range right_diags_reversed {
		right_diag_sum += count_xmas(s)
	}
	fmt.Println("--- RIGHT DIAGS ---")
	fmt.Println(right_diag_sum)
	//
	left_diags := append(get_diag(raw_split_rev, 0, 9), get_diag(raw_split_rev, 9, 1)...)
	// fmt.Println(get_diag(raw_split_rev, 0, 9))
	// fmt.Println(get_diag(raw_split_rev, 9, 1))
	// fmt.Println(left_diags)
	left_diags_reversed := make([]string, len(left_diags))
	for i, s := range left_diags {
		split := strings.Split(s, "")
		slices.Reverse(split)
		left_diags_reversed[i] = strings.Join(split, "")
	}
	// now find XMAS in every item in the slices
	left_diag_sum := 0
	for _, s := range left_diags {
		left_diag_sum += count_xmas(s)
	}
	for _, s := range left_diags_reversed {
		left_diag_sum += count_xmas(s)
	}
	fmt.Println("--- LEFT DIAGS ---")
	fmt.Println(left_diag_sum)
	fmt.Println(raw_split_rev)
}
