package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type hand struct {
	cards  []int
	bid    int
	counts map[int]int
	rank   int
}

type hands []hand

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func process_hand(hand []string) []int {
	// Map cards to values
	card_values := map[string]int{
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"T": 10,
		"J": 11,
		"Q": 12,
		"K": 13,
		"A": 14,
	}
	// Create slice of card values
	card_values_slice := []int{}
	for _, card := range hand {
		card_values_slice = append(card_values_slice, card_values[card])
	}
	return card_values_slice
}

func count_cards(hand []int) map[int]int {
	counts := make(map[int]int)

	for _, card := range hand {
		counts[card]++
	}

	return counts
}

func five_of_a_kind(hand hand) bool {
	for _, count := range hand.counts {
		if count == 5 {
			return true
		}
	}
	return false
}

func four_of_a_kind(hand hand) bool {
	for _, count := range hand.counts {
		if count == 4 {
			return true
		}
	}
	return false
}

func full_house(hand hand) bool {
	if len(hand.counts) == 2 {
		return true
	}
	return false
}

func two_pair(hand hand) bool {
	if len(hand.counts) == 3 {
		return true
	}
	return false
}

func one_pair(hand hand) bool {
	if len(hand.counts) == 4 {
		return true
	}
	return false
}

func high_card(hand hand) bool {
	if len(hand.counts) == 5 {
		return true
	}
	return false
}

func rate_hand(hand hand) int {
	return 0
}

func part1() {
	// Read the whole file in
	content, err := os.ReadFile("test_input.txt")
	check(err)
	// Split on line break and ignore blank lines
	chunks := strings.Split(string(content), "\n")
	// Remove last element (empty line)
	chunks = chunks[:len(chunks)-1]

	// Process all card hands
	hands := hands{}
	for _, chunk := range chunks {
		// Split on space
		data := strings.Split(chunk, " ")
		// Process cards
		cards_str := strings.Split(data[0], "")
		// Process bid
		bid, err := strconv.Atoi(data[1])
		check(err)
		// Process hand
		cards := process_hand(cards_str)

		// Create hand
		hand := hand{cards: cards, bid: bid, counts: count_cards(cards), rank: 0}
		hands = append(hands, hand)
	}
	fmt.Println(hands)
}

func part2() {

}

func main() {
	part1()
	part2()
}
