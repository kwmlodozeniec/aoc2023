package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type hand struct {
	cards     []int
	bid       int
	counts    []int
	hand_type int
}

type hands []hand

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func process_hand(hand []string, joker_mode bool) []int {
	// Map cards to values
	card_values := make(map[string]int)
	if joker_mode {
		card_values["J"] = 2
		card_values["2"] = 3
		card_values["3"] = 4
		card_values["4"] = 5
		card_values["5"] = 6
		card_values["6"] = 7
		card_values["7"] = 8
		card_values["8"] = 9
		card_values["9"] = 10
		card_values["T"] = 11
		card_values["Q"] = 12
		card_values["K"] = 13
		card_values["A"] = 14
	} else {
		card_values["2"] = 2
		card_values["3"] = 3
		card_values["4"] = 4
		card_values["5"] = 5
		card_values["6"] = 6
		card_values["7"] = 7
		card_values["8"] = 8
		card_values["9"] = 9
		card_values["T"] = 10
		card_values["J"] = 11
		card_values["Q"] = 12
		card_values["K"] = 13
		card_values["A"] = 14
	}
	// Create slice of card values
	card_values_slice := []int{}
	for _, card := range hand {
		card_values_slice = append(card_values_slice, card_values[card])
	}
	return card_values_slice
}

func compare_hands(left_hand hand, right_hand hand) bool {
	// left hand is worse than right hand based on type
	if left_hand.hand_type != right_hand.hand_type {
		return left_hand.hand_type < right_hand.hand_type
	}

	left_cards := left_hand.cards
	right_cards := right_hand.cards

	for idx, left_card := range left_cards {
		right_card := right_cards[idx]
		if left_card != right_card {
			return left_card < right_card
		}
	}
	return false
}

// Sort interfaces
func (h hands) Len() int           { return len(h) }
func (h hands) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h hands) Less(i, j int) bool { return compare_hands(h[i], h[j]) }

func count_cards(hand []int) []int {
	counts_map := make(map[int]int)

	for _, card := range hand {
		counts_map[card]++
	}

	counts := []int{}

	for _, count := range counts_map {
		counts = append(counts, count)
	}

	return counts
}

func is_in_list(list []int, value int) (bool, int) {
	found := false
	occurrences := 0

	for _, item := range list {
		if item == value {
			found = true
			occurrences++
		}
	}
	return found, occurrences
}

func five_of_a_kind(hand hand, joker_count int) bool {
	found_5, _ := is_in_list(hand.counts, 5)
	found_4, _ := is_in_list(hand.counts, 4)
	found_3, _ := is_in_list(hand.counts, 3)
	found_2, _ := is_in_list(hand.counts, 2)
	return found_5 || found_4 && joker_count > 0 || found_3 && found_2 && joker_count > 0
}

func four_of_a_kind(hand hand, joker_count int) bool {
	found_4, _ := is_in_list(hand.counts, 4)
	found_3, _ := is_in_list(hand.counts, 3)
	found_2, occurrences_2 := is_in_list(hand.counts, 2)
	found_1, occurrences_1 := is_in_list(hand.counts, 1)
	return found_4 || found_3 && found_1 && occurrences_1 == 2 && joker_count > 0 || found_2 && occurrences_2 == 2 && found_1 && joker_count == 2
}

func three_of_a_kind(hand hand, joker_count int) bool {
	found_3, _ := is_in_list(hand.counts, 3)
	found_2, _ := is_in_list(hand.counts, 2)
	found_1, occurrences_1 := is_in_list(hand.counts, 1)
	return found_3 && found_1 && occurrences_1 == 2 || found_2 && found_1 && occurrences_1 == 3 && joker_count > 0
}

func full_house(hand hand, joker_count int) bool {
	found_3, _ := is_in_list(hand.counts, 3)
	found_2, occurrences_2 := is_in_list(hand.counts, 2)
	found_1, _ := is_in_list(hand.counts, 1)
	return found_2 && found_3 || found_2 && occurrences_2 == 2 && found_1 && joker_count == 1
}

func two_pair(hand hand, joker_count int) bool {
	found, occurrences := is_in_list(hand.counts, 2)
	return found && occurrences == 2
}

func one_pair(hand hand, joker_count int) bool {
	found_2, occurrences_2 := is_in_list(hand.counts, 2)
	found_1, occurrences_1 := is_in_list(hand.counts, 1)
	return found_2 && occurrences_2 == 1 && found_1 && occurrences_1 == 3 || found_1 && occurrences_1 == 5 && joker_count > 0
}

func high_card(hand hand, joker_count int) bool {
	return len(hand.counts) == 5
}

func get_hand_type(hand hand, joker_mode bool) int {
	joker_count := 0
	if joker_mode {
		_, count := is_in_list(hand.cards, 2) // J == 2
		joker_count = count
	}

	if five_of_a_kind(hand, joker_count) {
		return 7
	}
	if four_of_a_kind(hand, joker_count) {
		return 6
	}
	if full_house(hand, joker_count) {
		return 5
	}
	if three_of_a_kind(hand, joker_count) {
		return 4
	}
	if two_pair(hand, joker_count) {
		return 3
	}
	if one_pair(hand, joker_count) {
		return 2
	}
	if high_card(hand, joker_count) {
		return 1
	}
	return 0
}

func part1() {
	// Read the whole file in
	content, err := os.ReadFile("input.txt")
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
		cards := process_hand(cards_str, false)

		// Create hand
		hand := hand{cards: cards, bid: bid, counts: count_cards(cards), hand_type: 0}
		hands = append(hands, hand)
	}

	// Get hand type
	for idx, hand := range hands {
		hands[idx].hand_type = get_hand_type(hand, false)
	}

	// Sort hands
	sort.Sort(hands)

	// Calculate total winnings
	total_winnings := 0
	for idx, hand := range hands {
		total_winnings += hand.bid * (idx + 1)
	}
	fmt.Println("Part 1: ", total_winnings)
}

func part2() {
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
		cards := process_hand(cards_str, true)

		// Create hand
		hand := hand{cards: cards, bid: bid, counts: count_cards(cards), hand_type: 0}
		hands = append(hands, hand)
	}

	// Get hand type
	for idx, hand := range hands {
		hands[idx].hand_type = get_hand_type(hand, true)
	}

	// Sort hands
	sort.Sort(hands)

	// Calculate total winnings
	total_winnings := 0
	for idx, hand := range hands {
		total_winnings += hand.bid * (idx + 1)
	}
	fmt.Println("Part 2: ", total_winnings)
}

func main() {
	part1()
	part2()
}
