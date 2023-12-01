package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var digitWords = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

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
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		chars := []string{}
		for _, char := range line {
			if unicode.IsDigit(char) {
				chars = append(chars, string(char))
			}
		}
		number_str := chars[0] + chars[len(chars)-1]
		number, err := strconv.Atoi(number_str)
		check(err)
		sum += number
	}
	fmt.Println("Part 1: ", sum)
}

func findAllSubstringPositions(fullString, subString string) []int {
	positions := []int{}
	index := -1

	for {
		index = strings.Index(fullString, subString)
		if index == -1 {
			break
		}

		positions = append(positions, index)
		fullString = fullString[index+len(subString):]
	}

	return positions
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
		chars := []string{}

		for idx, char := range line {
			// if we have a digit we can continue to the next char
			if unicode.IsDigit(char) {
				chars = append(chars, string(char))
				continue
			}

			// iterate over the digit words
			for word, digit := range digitWords {
				if char == rune(word[0]) { // if the first letter matches
					// check if the word fits in the line and is in the right place
					if len(word)+idx <= len(line) && word == line[idx:len(word)+idx] {
						chars = append(chars, digit)
						break
					}
				}
			}
		}
		number_str := chars[0] + chars[len(chars)-1]
		number, err := strconv.Atoi(number_str)
		check(err)
		sum += number
	}
	fmt.Println("Part 2: ", sum)
}

func main() {
	part1()
	part2()
}
