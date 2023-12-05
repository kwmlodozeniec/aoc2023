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
	number            int
	is_winner         bool
	winning_numbers   []int
	scratched_numbers []int
	matches           int
	score             int
	copies            int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func extract_numbers(line string) []int {
	numbers_list := []int{}
	numbers := regexp.MustCompile(`\b\d+\b`).FindAllString(line, -1)
	for _, number := range numbers {
		num, err := strconv.Atoi(number)
		check(err)
		numbers_list = append(numbers_list, num)
	}
	return numbers_list
}

func check_card_is_winner(card card) (bool, int) {
	found_numbers := 0
	for _, scratched_number := range card.scratched_numbers {
		for _, winning_number := range card.winning_numbers {
			if scratched_number == winning_number {
				found_numbers++
			}
		}
	}

	is_winner := bool(found_numbers > 0)
	return is_winner, found_numbers
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
		winning_numbers_data := extract_numbers(numbers_data[0])
		scratched_numbers := extract_numbers(numbers_data[1])
		card_number := extract_numbers(card_data[0])[0]
		check(err)

		// Create a card
		card := card{
			number:            card_number,
			is_winner:         false,
			winning_numbers:   winning_numbers_data,
			scratched_numbers: scratched_numbers,
		}
		is_winner, found_numbers := check_card_is_winner(card)
		card.is_winner = is_winner
		card.matches = found_numbers

		cards = append(cards, card)
	}

	// Calculate scores for all winning cards and total score
	total_score := 0
	for _, card := range cards {
		if card.is_winner {
			card.score = int(math.Pow(2.0, float64(card.matches-1)))
			total_score += card.score
		}
	}

	fmt.Println("Part 1: ", total_score)
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
		winning_numbers_data := extract_numbers(numbers_data[0])
		scratched_numbers := extract_numbers(numbers_data[1])
		card_number := extract_numbers(card_data[0])[0]
		check(err)

		// Create a card
		card := card{
			number:            card_number,
			is_winner:         false,
			winning_numbers:   winning_numbers_data,
			scratched_numbers: scratched_numbers,
			copies:            1,
		}
		is_winner, found_numbers := check_card_is_winner(card)
		card.is_winner = is_winner
		card.matches = found_numbers

		cards = append(cards, card)
	}

	// Process all cards
	for idx1, card := range cards {
		for idx2 := 0; idx2 < card.copies; idx2++ {
			for idx3 := idx1 + 1; idx3 <= idx1+card.matches; idx3++ {
				cards[idx3].copies++
			}
		}
	}

	// Calculate total cards
	total_cards := 0
	for _, card := range cards {
		total_cards += card.copies
	}

	fmt.Println("Part 2: ", total_cards)
}

func main() {
	part1()
	part2()
}
