package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type race struct {
	time_allowance     int
	best_distance      int
	distance_travelled []int
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
	content, err := os.ReadFile("input.txt")
	check(err)
	// Split on double line break
	chunks := strings.Split(string(content), "\n")

	// Extract the time and distance data
	time_data := chunks[0]
	distance_data := chunks[1]
	times := find_numbers(time_data)
	distances := find_numbers(distance_data)

	// "zip" the two slices together
	races := []race{}
	for idx := range times {
		races = append(races, race{times[idx], distances[idx], []int{}})
	}

	// Iterate over races and populate distance travelled for all times up to the time allowance
	for idx, race := range races {
		for i := 0; i < race.time_allowance; i++ {
			speed := 0 + i
			time_left := race.time_allowance - i
			distance_travelled := time_left * speed
			races[idx].distance_travelled = append(races[idx].distance_travelled, distance_travelled)
		}
	}

	// Find the hold times that beat the best distance
	hold_time_counts := []int{}
	for _, race := range races {
		hold_time_count := 0
		for _, distance := range race.distance_travelled {
			if distance > race.best_distance {
				hold_time_count++
			}
		}
		hold_time_counts = append(hold_time_counts, hold_time_count)
	}

	// Calculate the product of the hold times
	product := 1
	for _, hold_time_count := range hold_time_counts {
		product *= hold_time_count
	}

	fmt.Println("Part 1: ", product)
}

func part2() {
	// Read the whole file in
	content, err := os.ReadFile("input.txt")
	check(err)
	// Split on double line break
	chunks := strings.Split(string(content), "\n")

	// Extract the time and distance data
	time_data := chunks[0]
	distance_data := chunks[1]
	times := find_numbers(time_data)
	distances := find_numbers(distance_data)

	time_string := []string{}
	for _, time := range times {
		time_string = append(time_string, strconv.Itoa(time))
	}

	distance_string := []string{}
	for _, distance := range distances {
		distance_string = append(distance_string, strconv.Itoa(distance))
	}

	time, err := strconv.Atoi(strings.Join(time_string, ""))
	check(err)
	distance, err := strconv.Atoi(strings.Join(distance_string, ""))
	check(err)

	race := race{
		time_allowance:     time,
		best_distance:      distance,
		distance_travelled: []int{},
	}

	// Iterate over race and populate distance travelled for all times up to the time allowance
	for i := 0; i < race.time_allowance; i++ {
		speed := 0 + i
		time_left := race.time_allowance - i
		distance_travelled := time_left * speed
		race.distance_travelled = append(race.distance_travelled, distance_travelled)
	}

	// Find the hold times that beat the best distance
	hold_time_count := 0
	for _, distance := range race.distance_travelled {
		if distance > race.best_distance {
			hold_time_count++
		}
	}

	fmt.Println("Part 2: ", hold_time_count)
}

func main() {
	part1()
	part2()
}
