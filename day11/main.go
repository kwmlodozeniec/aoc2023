package main

import (
	"fmt"
	"os"
	"strings"
)

type point struct {
	row, col int
}

type galaxy struct {
	name int
	loc  point
}

func get_lines(file string) ([]string, []int, []int) {
	data, _ := os.ReadFile(file)
	lines := strings.Split(string(data), "\n")
	lines = lines[:len(lines)-1]
	empty_rows := get_empty_rows(lines)
	empty_cols := get_empty_cols(lines)
	return lines, empty_rows, empty_cols
}

func get_empty_rows(grid []string) []int {
	empty_rows := []int{}

	for row, line := range grid {
		empty_row := true
		for _, char := range line {
			if char != '.' {
				empty_row = false
				break
			}
		}
		if empty_row {
			empty_rows = append(empty_rows, row)
		}
	}
	return empty_rows
}

func get_empty_cols(grid []string) []int {
	empty_cols := []int{}

	for col := 0; col < len(grid[0]); col++ {
		empty_col := true
		for _, line := range grid {
			if line[col] != '.' {
				empty_col = false
				break
			}
		}
		if empty_col {
			empty_cols = append(empty_cols, col)
		}
	}
	return empty_cols
}

func get_galaxy_locations(grid []string) []galaxy {
	galaxies := []galaxy{}

	galaxy_count := 0
	for row, line := range grid {
		for col, char := range line {
			if char == '#' {
				galaxy_count++
				galaxies = append(galaxies, galaxy{galaxy_count, point{row, col}})
			}
		}
	}
	return galaxies
}

func is_in_list(list []int, number int) bool {
	for _, item := range list {
		if item == number {
			return true
		}
	}
	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func part1(input_file string) {
	grid, empty_rows, empty_cols := get_lines(input_file)

	galaxies := get_galaxy_locations(grid)

	expansion_factor := 2
	total := 0

	for idx, point := range galaxies {
		for _, to_the_left := range galaxies[:idx] {
			// get min/max in case numbers are out of order
			row_min := min(point.loc.row, to_the_left.loc.row)
			row_max := max(point.loc.row, to_the_left.loc.row)
			col_min := min(point.loc.col, to_the_left.loc.col)
			col_max := max(point.loc.col, to_the_left.loc.col)

			for r := row_min; r < row_max; r++ {
				if is_in_list(empty_rows, r) {
					total += expansion_factor
				} else {
					total++
				}
			}
			for c := col_min; c < col_max; c++ {
				if is_in_list(empty_cols, c) {
					total += expansion_factor
				} else {
					total++
				}
			}
		}
	}

	fmt.Println("Part 1:", total)
}

func part2(input_file string) {
	grid, empty_rows, empty_cols := get_lines(input_file)

	galaxies := get_galaxy_locations(grid)

	expansion_factor := 1000000
	total := 0

	for idx, point := range galaxies {
		for _, to_the_left := range galaxies[:idx] {
			// get min/max in case numbers are out of order
			row_min := min(point.loc.row, to_the_left.loc.row)
			row_max := max(point.loc.row, to_the_left.loc.row)
			col_min := min(point.loc.col, to_the_left.loc.col)
			col_max := max(point.loc.col, to_the_left.loc.col)

			for r := row_min; r < row_max; r++ {
				if is_in_list(empty_rows, r) {
					total += expansion_factor
				} else {
					total++
				}
			}
			for c := col_min; c < col_max; c++ {
				if is_in_list(empty_cols, c) {
					total += expansion_factor
				} else {
					total++
				}
			}
		}
	}

	fmt.Println("Part 2:", total)
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
