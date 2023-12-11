package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func find_numbers(line string) []int {
	str_fields := strings.Fields(strings.TrimSpace(line))
	found_numbers := make([]int, len(str_fields))
	for idx, str_field := range str_fields {
		number, err := strconv.Atoi(str_field)
		check(err)
		found_numbers[idx] = number
	}

	return found_numbers
}

func get_min(numbers []int) int {
	min := numbers[0]
	for _, number := range numbers {
		if number < min {
			min = number
		}
	}
	return min
}

func get_max(numbers []int) int {
	max := numbers[0]
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max
}

func sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func is_all_zeros(numbers []int) bool {
	if get_min(numbers) == 0 && get_max(numbers) == 0 {
		return true
	}
	return false
}

func get_diff_array(numbers []int) []int {
	if len(numbers) < 2 {
		return []int{0}
	}
	diff_array := []int{}

	for idx := len(numbers) - 1; idx >= 0; idx-- {
		if idx == 0 {
			continue
		}
		difference := numbers[idx] - numbers[idx-1]
		diff_array = append([]int{difference}, diff_array...)
	}
	return diff_array
}

func part1(input_file string) {
	// Read the whole file in
	content, err := os.ReadFile(input_file)
	check(err)
	// Split on line break and ignore blank lines
	chunks := strings.Split(string(content), "\n")
	// Remove last element (empty line)
	chunks = chunks[:len(chunks)-1]

	sequences := [][]int{}
	for _, chunk := range chunks {
		numbers := find_numbers(chunk)
		sequences = append(sequences, numbers)
	}

	final_numbers := []int{}
	for seq_idx, sequence := range sequences {
		// Calculate diff arrays
		diff_array := sequence
		diff_arrays := [][]int{}
		for !is_all_zeros(diff_array) {
			if len(diff_array) == 1 {
				break
			}
			new_diff_array := get_diff_array(diff_array)
			diff_arrays = append(diff_arrays, new_diff_array)
			diff_array = new_diff_array
		}

		// Figure out last element of each diff array
		// Work backwards
		for idx := len(diff_arrays) - 1; idx >= 0; idx-- {
			if is_all_zeros(diff_arrays[idx]) {
				diff_arrays[idx] = append(diff_arrays[idx], 0)
				continue
			}

			prev_array := diff_arrays[idx+1]
			prev_array_len := len(prev_array)
			last_element_prev_array := prev_array[prev_array_len-1]
			last_element_diff_array := diff_arrays[idx][len(diff_arrays[idx])-1]
			diff_arrays[idx] = append(diff_arrays[idx], last_element_prev_array+last_element_diff_array)
		}

		// Get last number in first diff array
		last_element := diff_arrays[0][len(diff_arrays[0])-1]
		// Add that number to the last number in the sequence
		final_number := sequence[len(sequence)-1] + last_element
		// Append to sequence
		sequences[seq_idx] = append(sequences[seq_idx], final_number)
		final_numbers = append(final_numbers, final_number)
	}
	fmt.Println("Part 1: ", sum(final_numbers))
}

func part2(input_file string) {
	// Read the whole file in
	content, err := os.ReadFile(input_file)
	check(err)
	// Split on line break and ignore blank lines
	chunks := strings.Split(string(content), "\n")
	// Remove last element (empty line)
	chunks = chunks[:len(chunks)-1]

	sequences := [][]int{}
	for _, chunk := range chunks {
		numbers := find_numbers(chunk)
		sequences = append(sequences, numbers)
	}

	final_numbers := []int{}
	for seq_idx, sequence := range sequences {
		// Calculate diff arrays
		diff_array := sequence
		diff_arrays := [][]int{}
		for !is_all_zeros(diff_array) {
			if len(diff_array) == 1 {
				break
			}
			new_diff_array := get_diff_array(diff_array)
			diff_arrays = append(diff_arrays, new_diff_array)
			diff_array = new_diff_array
		}

		// Figure out first element of each diff array
		// Work backwards
		for idx := len(diff_arrays) - 1; idx >= 0; idx-- {
			if is_all_zeros(diff_arrays[idx]) {
				diff_arrays[idx] = append([]int{0}, diff_arrays[idx]...)
				continue
			}

			first_element_prev_array := diff_arrays[idx+1][0]
			first_element_diff_array := diff_arrays[idx][0]
			diff_arrays[idx] = append([]int{first_element_diff_array - first_element_prev_array}, diff_arrays[idx]...)
		}

		// Get last first in first diff array
		first_element := diff_arrays[0][0]
		// Take away that number from the first number in the sequence
		final_number := sequence[0] - first_element
		// Append to sequence
		sequences[seq_idx] = append([]int{final_number}, sequences[seq_idx]...)
		final_numbers = append([]int{final_number}, final_numbers...)
	}
	fmt.Println("Part 2: ", sum(final_numbers))
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing argument: input file")
		return
	}
	input_file := os.Args[1]
	part1(input_file)
	part2(input_file)
}
