package day4

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Solve(input string) int {
	// Create a variable to keep track of the score for all cards.
	var totalScore = 0

	// Split the input into distinct cards by line.
	cards := strings.Split(input, "\n")

	// Split the cards into their respective sets (winners and numbers we have).
	for _, card := range cards {
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
					if score == 0 {
						score = 1
					} else {
						score *= 2
					}
				}
			}
		}

		totalScore += score
	}

	return totalScore
}
