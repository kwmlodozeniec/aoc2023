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
}

func main() {
	part1()
	part2()
}
