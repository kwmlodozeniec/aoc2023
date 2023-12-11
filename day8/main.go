package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type node map[string]string

type nodes map[string]node

func get_keys_that_end_with_char(nodes nodes, char string) []string {
	result := make(map[string]node)
	for key, value := range nodes {
		chars := strings.Split(key, "")
		if chars[len(chars)-1] == char {
			result[key] = value
		}
	}

	keys := []string{}
	for k := range result {
		keys = append(keys, k)
	}
	return keys
}

func all_end_with_char(nodes []string, char string) bool {
	for _, node := range nodes {
		chars := strings.Split(node, "")
		if len(chars) > 0 && chars[len(chars)-1] != char {
			return false
		}
	}
	return true
}

func get_min(data []int) int {
	element := data[0]
	for _, value := range data {
		if value < element {
			element = value
		}
	}
	return element
}

// Calculate greatest common divisor of two integers
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// calculate least common multiple of two integers
func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

// Calculate LCMs for an array
func lcms(numbers []int) int {
	result := numbers[0]
	for _, number := range numbers {
		result = lcm(result, number)
	}
	return result
}

func part1() {
	// Read the whole file in
	content, err := os.ReadFile("input.txt")
	check(err)
	// Split on line break and ignore blank lines
	chunks := strings.Split(string(content), "\n\n")

	// Extract instructions
	instr := strings.Split(chunks[0], "")

	// Extract node values
	node_values := strings.Split(chunks[1], "\n")
	// Remove last blank line
	node_values = node_values[:len(node_values)-1]

	// Extract nodes
	nodes := make(nodes)
	for _, node_line := range node_values {
		node_name := strings.Split(node_line, " = ")[0]
		node_values_str := strings.Split(node_line, " = ")[1]
		node_values_str = strings.ReplaceAll(node_values_str, "(", "")
		node_values_str = strings.ReplaceAll(node_values_str, ")", "")
		node_values := strings.Split(node_values_str, ", ")
		nodes[node_name] = node{
			"L": node_values[0],
			"R": node_values[1],
		}
	}

	// Find the path
	steps := 0
	current_node := "AAA"
	for current_node != "ZZZ" {
		for _, direction := range instr {
			current_node = nodes[current_node][direction]
			steps++
		}
	}

	fmt.Println("Part 1: ", steps)
}

func part2() {
	// Read the whole file in
	content, err := os.ReadFile("input.txt")
	check(err)
	// Split on line break and ignore blank lines
	chunks := strings.Split(string(content), "\n\n")

	// Extract instructions
	instr := strings.Split(chunks[0], "")

	// Extract node values
	node_values := strings.Split(chunks[1], "\n")
	// Remove last blank line
	node_values = node_values[:len(node_values)-1]

	// Extract nodes
	nodes := make(nodes)
	for _, node_line := range node_values {
		node_name := strings.Split(node_line, " = ")[0]
		node_values_str := strings.Split(node_line, " = ")[1]
		node_values_str = strings.ReplaceAll(node_values_str, "(", "")
		node_values_str = strings.ReplaceAll(node_values_str, ")", "")
		node_values := strings.Split(node_values_str, ", ")
		nodes[node_name] = node{
			"L": node_values[0],
			"R": node_values[1],
		}
	}

	// Find the path
	steps := 0
	current_nodes := get_keys_that_end_with_char(nodes, "A")
	end_found := make([]int, len(current_nodes))

	// fmt.Println("Current nodes: ", current_nodes)
	for get_min(end_found) == 0 {
		for _, direction := range instr {
			for idx, current_node := range current_nodes {
				if end_found[idx] == 0 {
					// fmt.Println("We are at step ", steps, " and node ", idx, " is ", current_node, " and direction is ", direction)
					current_nodes[idx] = nodes[current_node][direction]
					if current_nodes[idx][len(current_nodes[idx])-1:] == "Z" {
						end_found[idx] = steps + 1
						// fmt.Println("Found end for node ", idx, " at step ", steps+1)
					}
				} else {
					// fmt.Println("Node ", idx, " already found")
				}
			}
			steps++
		}
		// fmt.Println("Current nodes: ", current_nodes, " and end_found: ", end_found)
	}

	fmt.Println("Part 2: ", lcms(end_found))
}

func main() {
	part1()
	part2()
}
