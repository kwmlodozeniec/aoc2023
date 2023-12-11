package main

import (
	"fmt"
	"os"
	"strings"
)

type point struct {
	x, y int
}

func get_lines(file string) []string {
	data, _ := os.ReadFile(file)
	return strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
}

func find_start(grid []string) point {
	for y, line := range grid {
		for x, char := range line {
			if char == 'S' {
				return point{x, y}
			}
		}
	}
	return point{-1, -1}
}

func get_max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func get_next_points(grid []string, current_point point) []point {
	next_points := []point{}

	switch grid[current_point.y][current_point.x] {
	case '|': // up or down
		next_points = append(next_points, point{current_point.x, current_point.y - 1})
		next_points = append(next_points, point{current_point.x, current_point.y + 1})
	case '-': // left or right
		next_points = append(next_points, point{current_point.x - 1, current_point.y})
		next_points = append(next_points, point{current_point.x + 1, current_point.y})
	case 'L': // up and right
		next_points = append(next_points, point{current_point.x, current_point.y - 1})
		next_points = append(next_points, point{current_point.x + 1, current_point.y})
	case 'J': // up and left
		next_points = append(next_points, point{current_point.x, current_point.y - 1})
		next_points = append(next_points, point{current_point.x - 1, current_point.y})
	case '7': // down and left
		next_points = append(next_points, point{current_point.x, current_point.y + 1})
		next_points = append(next_points, point{current_point.x - 1, current_point.y})
	case 'F': // down and right
		next_points = append(next_points, point{current_point.x, current_point.y + 1})
		next_points = append(next_points, point{current_point.x + 1, current_point.y})
	case '.':
	case 'S':
		// get surrounding points, account for points on the edge of the grid
		up := byte(' ')
		if current_point.y > 0 {
			up = grid[current_point.y-1][current_point.x]
		}

		down := byte(' ')
		if current_point.y < len(grid)-1 {
			down = grid[current_point.y+1][current_point.x]
		}

		left := byte(' ')
		if current_point.x > 0 {
			left = grid[current_point.y][current_point.x-1]
		}

		right := byte(' ')
		if current_point.x < len(grid[current_point.y])-1 {
			right = grid[current_point.y][current_point.x+1]
		}

		// points above are only valid if we can go straight down, or from the left or right
		if up == '|' || up == '7' || up == 'F' {
			next_points = append(next_points, point{current_point.x, current_point.y - 1})
		}
		// points below are only valid if we can go straight up, or from the left or right
		if down == '|' || down == 'J' || down == 'L' {
			next_points = append(next_points, point{current_point.x, current_point.y + 1})
		}
		// points to the left are only valid if we can go straight, or from the right
		if left == '-' || left == 'L' || left == 'F' {
			next_points = append(next_points, point{current_point.x - 1, current_point.y})
		}
		// points ro the right are only valid if we can go straight, or from the left
		if right == '-' || right == 'J' || right == '7' {
			next_points = append(next_points, point{current_point.x + 1, current_point.y})
		}
	}
	return next_points
}

func part1(input_file string) {
	grid := get_lines(input_file)
	start := find_start(grid)
	// Keep track of visited points
	visited := map[point]int{start: 0}
	points_to_check := []point{start}

	// keep track of distance from start
	distance := 0
	for len(points_to_check) > 0 {
		current_point := points_to_check[0]
		// remove current point from points to check
		points_to_check = points_to_check[1:]
		// get next points to visit
		next_points := get_next_points(grid, current_point)

		// visit next points
		for _, point := range next_points {
			// if we haven't visited this point yet
			if _, ok := visited[point]; !ok {
				// mark as visited and add to its distance from the previous point
				visited[point] = visited[current_point] + 1
				// update max distance
				distance = get_max(distance, visited[current_point]+1)
				// add to points to check
				points_to_check = append(points_to_check, point)
			}
		}
	}
	fmt.Println("Part 1: ", distance)

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
