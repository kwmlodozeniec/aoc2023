package main

import (
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
	destination_start int
	source_start      int
	length            int
}

type seed_to_soil generic_map
type fertilizer_to_water generic_map
type water_to_light generic_map
type light_to_temp generic_map
type temp_to_humidity generic_map
type humidity_to_location generic_map

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func find_numbers(line string) ([]int, error) {
	found_numbers := []int{}

	numbers := regexp.MustCompile(`\b\d+\b`).FindAllStringIndex(line, -1)
	for _, number := range numbers {
		value, err := strconv.Atoi(line[number[0]:number[1]])
		check(err)
		found_numbers = append(found_numbers, value)
	}

	return found_numbers, nil
}

func part1() {
	// Read the whole file in
	content, err := os.ReadFile("test_input.txt")
	check(err)
	// Split on double line break
	chunks := strings.Split(string(content), "\n\n")

	seeds_line := chunks[0]
	map_data_blocks := chunks[1:]

	// Iterate over the maps
	seed_to_soil_maps := []seed_to_soil{}
	fertilizer_to_water_maps := []fertilizer_to_water{}
	water_to_light_maps := []water_to_light{}
	light_to_temp_maps := []light_to_temp{}
	temp_to_humidity_maps := []temp_to_humidity{}
	humidly_to_location_maps := []humidity_to_location{}

	for _, map_data := range map_data_blocks {
		lines := strings.Split(map_data, "\n")
	}
}

func part2() {
}

func main() {
	part1()
	part2()
}
