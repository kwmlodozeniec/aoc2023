package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type partNumber struct {
	line     int
	startIdx int
	endIdx   int
	value    int
	isValid  bool
}

type symbol struct {
	line      int
	idx       int
	value     string
	isGear    bool
	gearSizeA int
	gearSizeB int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func findNumbers(line string, lineNumber int) ([]partNumber, error) {
	partNumbers := []partNumber{}

	numbers := regexp.MustCompile(`\b\d+\b`).FindAllStringIndex(line, -1)
	for _, number := range numbers {
		fmt.Println("Line: ", line, "Number start/ends: ", number)
		partNumber := partNumber{
			line:     lineNumber,
			startIdx: number[0],
			endIdx:   number[1] - 1,
			value:    0,
			isValid:  false,
		}
		value, err := strconv.Atoi(line[partNumber.startIdx : partNumber.endIdx+1])
		check(err)
		partNumber.value = value
		partNumbers = append(partNumbers, partNumber)
	}

	return partNumbers, nil
}

func findSymbols(line string, lineNumber int) ([]symbol, error) {
	symbols := []symbol{}

	foundSymbols := regexp.MustCompile(`[^\w\d.\n\r]`).FindAllStringIndex(line, -1)
	for _, foundSymbol := range foundSymbols {
		fmt.Println("Line: ", line, "Symbol index: ", foundSymbol)
		symbol := symbol{
			line:  lineNumber,
			idx:   foundSymbol[0],
			value: line[foundSymbol[0]:foundSymbol[1]],
		}
		symbols = append(symbols, symbol)
	}

	return symbols, nil
}

func validatePartNumber(partNumber partNumber, symbols []symbol) (partNumber, error) {
	for _, symbol := range symbols {
		if symbol.line >= partNumber.line-1 && // Check if symbol is on the line before or on the same line
			symbol.line <= partNumber.line+1 && // Check if symbol is on the line after or on the same line
			symbol.idx >= partNumber.startIdx-1 && // Check if symbol index overlaps with the part number from the left
			symbol.idx <= partNumber.endIdx+1 { // Check if symbol index overlaps with the part number from the right  {
			partNumber.isValid = true
			fmt.Println("Part number: ", partNumber.value, "is valid")
		}
	}
	return partNumber, nil
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
	allPartNumbers := []partNumber{}
	allSymbols := []symbol{}
	for scanner.Scan() {
		numbers, err := findNumbers(scanner.Text(), line)
		check(err)
		symbols, err := findSymbols(scanner.Text(), line)
		check(err)
		allPartNumbers = append(allPartNumbers, numbers...)
		allSymbols = append(allSymbols, symbols...)
		line++
	}
	fmt.Println(allPartNumbers)
	fmt.Println(allSymbols)

	// Validate part numbers
	validatedPartNumbers := []partNumber{}
	for _, part := range allPartNumbers {
		part, err := validatePartNumber(part, allSymbols)
		check(err)
		validatedPartNumbers = append(validatedPartNumbers, part)
	}

	sum := 0
	for _, part := range validatedPartNumbers {
		if part.isValid {
			sum += part.value
		}
	}
	fmt.Println("Part 1: ", sum)
}

func part2() {
}

func main() {
	part1()
	part2()
}
