package main

import (
	"bufio"
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

func part1() {
	// Open file
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Iterate over the lines
	colour_limits := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		game_data := strings.Split(line, ": ")
		game_id_data, set_data := game_data[0], game_data[1]
		game_id, err := strconv.Atoi(strings.Split(game_id_data, " ")[1])
		check(err)
		draws := strings.Split(set_data, "; ")
		game_valid := true
		for _, draw := range draws {
			colours := strings.Split(draw, ", ")
			for _, colour_count := range colours {
				count_colour := strings.Split(colour_count, " ")
				ball_count, err := strconv.Atoi(count_colour[0])
				check(err)
				ball_colour := count_colour[1]

				if ball_count > colour_limits[ball_colour] {
					game_valid = false
					break
				}
			}
		}
		if game_valid {
			sum += game_id
		}
	}
	fmt.Println("Part 1: ", sum)
}

func part2() {
	// Open file
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Iterate over the lines
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		game_data := strings.Split(line, ": ")
		set_data := game_data[1]
		check(err)
		draws := strings.Split(set_data, "; ")

		colour_counts := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for _, draw := range draws {
			colours := strings.Split(draw, ", ")
			for _, colour_count := range colours {
				count_colour := strings.Split(colour_count, " ")
				ball_count, err := strconv.Atoi(count_colour[0])
				check(err)
				ball_colour := count_colour[1]
				if ball_count > colour_counts[ball_colour] {
					colour_counts[ball_colour] = ball_count
				}
			}
		}
		product := colour_counts["red"] * colour_counts["green"] * colour_counts["blue"]
		sum += product
	}
	fmt.Println("Part 2: ", sum)
}

func main() {
	part1()
	part2()
}
