package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type card struct {
	number           int
	isWinner         bool
	winningNumbers   []int
	scratchedNumbers []int
	matches          int
	score            int
	copies           int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func extractNumbers(line string) []int {
	numbersList := []int{}
	numbers := regexp.MustCompile(`\b\d+\b`).FindAllString(line, -1)
	for _, number := range numbers {
		num, err := strconv.Atoi(number)
		check(err)
		numbersList = append(numbersList, num)
	}
	return numbersList
}

func checkCardIsWinner(card card) (bool, int) {
	foundNumbers := 0
	for _, scratchedNumber := range card.scratchedNumbers {
		for _, winningNumber := range card.winningNumbers {
			if scratchedNumber == winningNumber {
				foundNumbers++
			}
		}
	}

	isWinner := bool(foundNumbers > 0)
	return isWinner, foundNumbers
}

func part1() {
	// Open file
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Iterate over the lines
	cards := []card{}
	for scanner.Scan() {
		// Get the line
		line := scanner.Text()
		card_data := strings.Split(line, ":")
		numbers_data := strings.Split(card_data[1], "|")
		winning_numbers_data := extractNumbers(numbers_data[0])
		scratched_numbers := extractNumbers(numbers_data[1])
		card_number := extractNumbers(card_data[0])[0]
		check(err)

		// Create a card
		card := card{
			number:           card_number,
			isWinner:         false,
			winningNumbers:   winning_numbers_data,
			scratchedNumbers: scratched_numbers,
		}
		isWinner, foundNumbers := checkCardIsWinner(card)
		card.isWinner = isWinner
		card.matches = foundNumbers

		cards = append(cards, card)
	}

	// Calculate scores for all winning cards and total score
	totalScore := 0
	for _, card := range cards {
		if card.isWinner {
			card.score = int(math.Pow(2.0, float64(card.matches-1)))
			totalScore += card.score
			// fmt.Println("Card ", card.number, " is a winner with ", card.matches, " matches and a score of ", card.score)
		}
	}

	fmt.Println("Part 1: ", totalScore)
}

func part2() {
	// Open file
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Iterate over the lines
	cards := []card{}
	for scanner.Scan() {
		// Get the line
		line := scanner.Text()
		card_data := strings.Split(line, ":")
		numbers_data := strings.Split(card_data[1], "|")
		winning_numbers_data := extractNumbers(numbers_data[0])
		scratched_numbers := extractNumbers(numbers_data[1])
		card_number := extractNumbers(card_data[0])[0]
		check(err)

		// Create a card
		card := card{
			number:           card_number,
			isWinner:         false,
			winningNumbers:   winning_numbers_data,
			scratchedNumbers: scratched_numbers,
			copies:           1,
		}
		isWinner, foundNumbers := checkCardIsWinner(card)
		card.isWinner = isWinner
		card.matches = foundNumbers

		cards = append(cards, card)
	}

	// Process all cards
	for idx1, card := range cards {
		// fmt.Println("Card ", card.number, "wins ", card.matches, "cards, and has ", card.copies, " copies")
		for idx2 := 0; idx2 < card.copies; idx2++ {
			for idx3 := idx1 + 1; idx3 <= idx1+card.matches; idx3++ {
				// fmt.Println("Increasing copies of card ", cards[idx3].number, " to ", cards[idx3].copies+1, " copies")
				cards[idx3].copies++
			}
		}
	}

	// Calculate total cards
	totalCards := 0
	for _, card := range cards {
		totalCards += card.copies
	}

	fmt.Println("Part 2: ", totalCards)
}

func main() {
	part1()
	part2()
}
