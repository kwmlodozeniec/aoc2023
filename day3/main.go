package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type part_number struct {
	line      int
	start_idx int
	end_idx   int
	value     int
	is_valid  bool
}

type symbol struct {
	line        int
	idx         int
	value       string
	is_gear     bool
	gear_size_a int
	gear_size_b int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func find_numbers(line string, line_number int) ([]part_number, error) {
	part_numbers := []part_number{}

	numbers := regexp.MustCompile(`\b\d+\b`).FindAllStringIndex(line, -1)
	for _, number := range numbers {
		part_number := part_number{
			line:      line_number,
			start_idx: number[0],
			end_idx:   number[1] - 1,
			value:     0,
			is_valid:  false,
		}
		value, err := strconv.Atoi(line[part_number.start_idx : part_number.end_idx+1])
		check(err)
		part_number.value = value
		part_numbers = append(part_numbers, part_number)
	}

	return part_numbers, nil
}

func find_symbols(line string, line_number int) ([]symbol, error) {
	symbols := []symbol{}

	found_symbols := regexp.MustCompile(`[^\w\d.\n\r]`).FindAllStringIndex(line, -1)
	for _, found_symbol := range found_symbols {
		symbol := symbol{
			line:  line_number,
			idx:   found_symbol[0],
			value: line[found_symbol[0]:found_symbol[1]],
		}
		symbols = append(symbols, symbol)
	}

	return symbols, nil
}

func validate_part_number(part part_number, symbols []symbol) (part_number, error) {
	for _, symbol := range symbols {
		if symbol.line >= part.line-1 && // Check if symbol is on the line before or on the same line
			symbol.line <= part.line+1 && // Check if symbol is on the line after or on the same line
			symbol.idx >= part.start_idx-1 && // Check if symbol index overlaps with the part number from the left
			symbol.idx <= part.end_idx+1 { // Check if symbol index overlaps with the part number from the right  {
			part.is_valid = true
		}
	}
	return part, nil
}

func validate_gears(part_numbers []part_number, symbols []symbol) ([]symbol, error) {
	gears := []symbol{}
	for _, symbol := range symbols {
		if symbol.value == "*" {
			surrounding_part_numbers := []part_number{}

			for _, part_number := range part_numbers {
				if part_number.line == symbol.line || part_number.line == symbol.line-1 || part_number.line == symbol.line+1 {
					surrounding_part_numbers = append(surrounding_part_numbers, part_number)
					// fmt.Println("Surrounding part number: ", partNumber.value, "Start index: ", partNumber.startIdx, "End index: ", partNumber.endIdx)
				}
			}

			adjacent_part_numbers := []part_number{}
			for _, part_number := range surrounding_part_numbers {
				if symbol.idx >= part_number.start_idx &&
					symbol.idx <= part_number.end_idx ||
					symbol.idx-part_number.start_idx == 1 ||
					part_number.start_idx-symbol.idx == 1 ||
					symbol.idx-part_number.end_idx == -1 ||
					part_number.end_idx-symbol.idx == -1 {
					adjacent_part_numbers = append(adjacent_part_numbers, part_number)
					// fmt.Println("Adjacent part number: ", partNumber.value)
				}
			}
			if len(adjacent_part_numbers) == 2 {
				symbol.is_gear = true
				symbol.gear_size_a = adjacent_part_numbers[0].value
				symbol.gear_size_b = adjacent_part_numbers[1].value
				gears = append(gears, symbol)
			}
		}
	}
	return gears, nil
}

func part1() {
	// Open file
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Iterate over the lines
	line := 0
	all_part_numbers := []part_number{}
	all_symbols := []symbol{}
	for scanner.Scan() {
		numbers, err := find_numbers(scanner.Text(), line)
		check(err)
		symbols, err := find_symbols(scanner.Text(), line)
		check(err)
		all_part_numbers = append(all_part_numbers, numbers...)
		all_symbols = append(all_symbols, symbols...)
		line++
	}

	// Validate part numbers
	validated_part_numbers := []part_number{}
	for _, part := range all_part_numbers {
		part, err := validate_part_number(part, all_symbols)
		check(err)
		validated_part_numbers = append(validated_part_numbers, part)
	}

	sum := 0
	for _, part := range validated_part_numbers {
		if part.is_valid {
			sum += part.value
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
	line := 0
	all_part_numbers := []part_number{}
	all_symbols := []symbol{}
	for scanner.Scan() {
		numbers, err := find_numbers(scanner.Text(), line)
		check(err)
		symbols, err := find_symbols(scanner.Text(), line)
		check(err)
		all_part_numbers = append(all_part_numbers, numbers...)
		all_symbols = append(all_symbols, symbols...)
		line++
	}

	// Validate part numbers
	validated_part_numbers := []part_number{}
	for _, part := range all_part_numbers {
		part, err := validate_part_number(part, all_symbols)
		check(err)
		validated_part_numbers = append(validated_part_numbers, part)
	}

	// Validate gears
	gears, err := validate_gears(validated_part_numbers, all_symbols)
	check(err)

	// Calculate sum of ratios
	sum := 0
	for _, gear := range gears {
		sum += gear.gear_size_a * gear.gear_size_b
	}
	fmt.Println("Part 2: ", sum)
}

func main() {
	part1()
	part2()
}
