package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type seed struct {
	number     int
	soil       int
	fertilizer int
	water      int
	light      int
	humidity   int
	location   int
}

type seeds []seed

type generic_map struct {
	name           string
	destination_start int
	source_start      int
	length            int
}

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

func part1() {
	// Read the whole file in
	content, err := os.ReadFile("test_input.txt")
	check(err)
	// Split on double line break
	chunks := strings.Split(string(content), "\n\n")

	seeds_line := chunks[0]
	map_data_blocks := chunks[1:]
	fmt.Println("Seeds line: ", seeds_line)
	fmt.Println("Map data blocks: ", map_data_blocks)

	// Parse the seeds
	seeds := find_numbers(seeds_line)
	fmt.Println("Seeds: ", seeds)

	map_list := map[int]string{
		0: "seed-to-soil",
		1: "soil-to-fertilizer",
		2: "fertilizer-to-water",
		3: "water-to-light",
		4: "light-to-temperature",
		5: "temperature-to-humidity",
		6: "humidity-to-location",
	}
	maps := []generic_map{}
	fmt.Println("Map list: ", map_list, " Maps: ", maps)

	// Build the maps
	for _, map_data := range map_data_blocks {
		lines := strings.Split(map_data, "\n")
		map_name := strings.Split(lines[0], " map:")[0]
		map_data_lines := lines[1:]
		fmt.Println("Map name: ", map_name, " Map data lines: ", map_data_lines)
	}
}

func part2() {
}

func main() {
	part1()
	part2()
}
