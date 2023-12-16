package day4

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Solve Day 4, returning a tuple of (part 1 total score, part 2 total cards).
func Solve(input string) (int, int) {
	// Create a variable to keep track of the score for all cards.
	var totalScore = 0

	// Split the input into distinct cards by line.
	cards := strings.Split(input, "\n")

	// For part 2, we'll need to keep track of the number of matches for each
	// card so we can calculate the total number of cards we'd have at the end.
	cardMatches := make([]int, len(cards))
	cardCount := make([]int, len(cards))

	// Split the cards into their respective sets (winners and numbers we have).
	for cardIdx, card := range cards {
		// Create a variable to keep track of the score for this one card.
		var score = 0

		// Parse the numbers (and the "|") from the card
		r := regexp.MustCompile(`(?:\w*\d+)+|\|`)
		symbols := r.FindAllStringSubmatch(card, -1)

		// We're going to need to distinguish between when we're parsing winners
		// vs. numbers we have, so a simple boolean works.
		var parsingWinners = true

		// And to match winners quickly, a set (you know, a map in go) will keep
		// track of the winners.
		// We don't think we need the value of these keys, but we can use them
		// to keep track of the number of winners that happen to be the same
		// number in case its actually important.
		winners := make(map[int]int)

		// For part 2, we'll need to keep track of how many matches we have and
		// what card number we're on. We can also go ahead and increase the card
		// count for this card.
		var matches = 0
		cardCount[cardIdx]++

		// Begin reading the parsed symbols, omitting the first because its the
		// card number.
		for _, symbol := range symbols[1:] {
			// We don't expect any symbols to have more than one match, so panic
			// on that shit.
			if len(symbol) > 1 {
				panic(fmt.Sprintf("Unexpected number of matches (%d) for symbol:%#v", len(symbol), symbol))
			}
			s := symbol[0]

			// If the symbol is the "|" character, we're switching from winners
			// to numbers we have.
			if s == "|" {
				parsingWinners = false
				continue
			}

			// Otherwise, game on. Parse the number.
			num, err := strconv.Atoi(s)
			if err != nil {
				panic(fmt.Sprintf("Expected to parse number from symbol: %s", s))
			}

			// If we're looking at winners, we need to add them to the set.
			if parsingWinners {
				winners[num] += 1
			} else {
				// Otherwise, we're scoring ourselves.
				if _, isWinner := winners[num]; isWinner {
					matches++
					if score == 0 {
						score = 1
					} else {
						score *= 2
					}
				}
			}
		}

		// For part 2, store the number of matches for this card, which we'll
		// use to calculate the total number of cards.
		cardMatches[cardIdx] = matches

		// Add this score to the total score for part 1
		totalScore += score
	}

	// To calculate the total number of cards, we first need to increase
	// the card counts to take into account duplicated cards from the previous
	// card's number of matching winners.
	for cardIdx, matches := range cardMatches {
		for i := 0; i < matches; i++ {
			cardCount[cardIdx+i+1] += cardCount[cardIdx]
		}
	}

	// Finally, sum the total number of cards.
	var totalCards = 0
	for _, cardCount := range cardCount {
		totalCards += cardCount
	}

	return totalScore, totalCards
}
