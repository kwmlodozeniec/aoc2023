package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type point struct {
	x, y int
}

type galaxy struct {
	name int
	loc  point
}

func get_lines(file string) []string {
	data, _ := os.ReadFile(file)
	lines := strings.Split(string(data), "\n")
	return lines[:len(lines)-1]
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
			fmt.Println("Empty row:", row)
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
			fmt.Println("Empty col:", col)
			empty_cols = append(empty_cols, col)
		}
	}
	return empty_cols
}

func expand_grid(grid []string) []string {
	empty_rows := get_empty_rows(grid)
	empty_cols := get_empty_cols(grid)

	// Add empty rows
	new_grid := []string{}
	for row_idx, row := range grid {
		new_grid = append(new_grid, row)
		if slices.Contains(empty_rows, row_idx) {
			new_grid = append(new_grid, row)
		}
	}

	// Add empty cols
	for row, line := range new_grid {
		new_line := strings.Split(line, "")
		for idx, col_idx := range empty_cols {
			new_line = append(new_line[:col_idx+idx], append([]string{"."}, new_line[col_idx+idx:]...)...)
		}
		new_grid[row] = strings.Join(new_line, "")
	}

	return new_grid
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

func gen_combinations(galaxies []galaxy, count int) [][]galaxy {
	result := [][]galaxy{}
	current := []galaxy{}

	var backtrack func(start int)
	backtrack = func(start int) {
		if len(current) == count {
			result = append(result, append([]galaxy{}, current...))
			return
		}
		for i := start; i < len(galaxies); i++ {
			current = append(current, galaxies[i])
			backtrack(i + 1)
			current = current[:len(current)-1]
		}
	}

	backtrack(0)
	return result
}

func part1(input_file string) {
	grid := get_lines(input_file)
	for idx, line := range grid {
		fmt.Printf("Row %2d: %s\n", idx, line)
	}
	grid = expand_grid(grid)
	for idx, line := range grid {
		fmt.Printf("Row %2d: %s\n", idx, line)
	}

	galaxies := get_galaxy_locations(grid)
	fmt.Println(galaxies)

	all_combinations := gen_combinations(galaxies, 2)
	for _, comb := range all_combinations {
		fmt.Println(comb)
	}
}

func part2(input_file string) {

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
