package main

import (
	"bufio"
	"os"
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
type humidly_to_location generic_map

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func part1() {
	// Open file
	file, err := os.Open("test_input.txt")
	check(err)
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Iterate over the lines
	cards := []card{}
	for scanner.Scan() {
		line := scanner.Text()
	}
}

func part2() {
}

func main() {
	part1()
	part2()
}
