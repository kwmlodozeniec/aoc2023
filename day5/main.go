package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func find_numbers(line string) []int {
	found_numbers := []int{}

	numbers := regexp.MustCompile(`\b\d+\b`).FindAllStringIndex(line, -1)
	for _, number := range numbers {
		value, err := strconv.Atoi(line[number[0]:number[1]])
		check(err)
		found_numbers = append(found_numbers, value)
	}

	return found_numbers
}

func min(numbers []int) int {
	min := numbers[0]
	for _, number := range numbers {
		if number < min {
			min = number
		}
	}
	return min
}

func max(numbers []int) int {
	max := numbers[0]
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max
}

func part1() {
	// Read the whole file in
	content, err := os.ReadFile("input.txt")
	check(err)
	// Split on double line break
	chunks := strings.Split(string(content), "\n\n")

	// Parse the seeds
	seeds := find_numbers(chunks[0])

	blocks := chunks[1:]

	// Parse the map data
	for _, block := range blocks {
		ranges := [][]int{}
		for _, line := range strings.Split(block, "\n")[1:] {
			if len(line) > 0 {
				ranges = append(ranges, find_numbers(line))
			}
		}
		new_locations := []int{}
		for _, seed := range seeds {
			found := false
			for _, r := range ranges {
				dest_start, seed_start, length := r[0], r[1], r[2]
				if seed_start <= seed && seed < seed_start+length {
					new_locations = append(new_locations, seed-seed_start+dest_start)
					found = true
					break
				}
			}
			if !found {
				new_locations = append(new_locations, seed)
			}
		}
		seeds = new_locations
	}
	fmt.Println("Part 1: ", min(seeds))
}

func part2() {
	// Read the whole file in
	content, err := os.ReadFile("input.txt")
	check(err)
	// Split on double line break
	chunks := strings.Split(string(content), "\n\n")

	blocks := chunks[1:]

	// Parse the seeds
	seed_range_data := find_numbers(chunks[0])
	seed_ranges := [][]int{}
	for i := 0; i < len(seed_range_data); i += 2 {
		seed_range := []int{seed_range_data[i], seed_range_data[i] + seed_range_data[i+1]}
		seed_ranges = append(seed_ranges, seed_range)
	}

	// Parse the map data
	for _, block := range blocks {
		ranges := [][]int{}
		for _, line := range strings.Split(block, "\n")[1:] {
			if len(line) > 0 {
				ranges = append(ranges, find_numbers(line))
			}
		}
		new_locations := [][]int{}
		for len(seed_ranges) > 0 {
			seed_range_start, seed_range_end := seed_ranges[len(seed_ranges)-1][0], seed_ranges[len(seed_ranges)-1][1]
			// pop the last elements of the slice
			seed_ranges = seed_ranges[:len(seed_ranges)-1]
			found := false
			for _, r := range ranges {
				dest_start, seed_start, length := r[0], r[1], r[2]
				overlap_start := max([]int{seed_range_start, seed_start})
				overlap_end := min([]int{seed_range_end, seed_start + length})

				if overlap_start < overlap_end {
					new_locations = append(new_locations, []int{overlap_start - seed_start + dest_start, overlap_end - seed_start + dest_start})
					if overlap_start > seed_range_start {
						seed_ranges = append(seed_ranges, []int{seed_range_start, overlap_start})
					}
					if seed_range_end > overlap_end {
						seed_ranges = append(seed_ranges, []int{overlap_end, seed_range_end})
					}
					found = true
					break
				}
			}
			if !found {
				new_locations = append(new_locations, []int{seed_range_start, seed_range_end})
			}

		}
		seed_ranges = new_locations
	}
	min_locations := []int{}
	for _, r := range seed_ranges {
		min_locations = append(min_locations, r[0])
	}
	fmt.Println("Part 2: ", min(min_locations))
}

func main() {
	part1()
	part2()
}
