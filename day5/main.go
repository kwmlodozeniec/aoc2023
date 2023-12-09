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
	name              string
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

	// Parse the seeds
	seeds := find_numbers(seeds_line)
	fmt.Println(seeds)

	map_list := map[int]string{
		0: "seed_to_soil",
		1: "soil_to_fertilizer",
		2: "fertilizer_to_water",
		3: "water_to_light",
		4: "light_to_temp",
		5: "temp_to_humidity",
		6: "humidity_to_location",
	}
	maps := []generic_map{}

	fmt.Println(map_list, maps)

	// Build the maps
	for _, map_data := range map_data_blocks {
		lines := strings.Split(map_data, "\n")[1:]

		for _, map_data := range lines {

		}
	}
}

func part2() {
}

func main() {
	part1()
	part2()
}
