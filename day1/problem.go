package day1

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

/*
*
Something is wrong with global snow production, and you've been selected to take a look. The Elves have even given you a map; on it, they've used stars to mark the top fifty locations that are likely to be having problems.

You've been doing this long enough to know that to restore snow operations, you need to check all fifty stars by December 25th.

Collect stars by solving puzzles. Two puzzles will be made available on each day in the Advent calendar; the second puzzle is unlocked when you complete the first. Each puzzle grants one star. Good luck!

You try to ask why they can't just use a weather machine ("not powerful enough") and where they're even sending you ("the sky") and why your map looks mostly blank ("you sure ask a lot of questions") and hang on did you just say the sky ("of course, where do you think snow comes from") when you realize that the Elves are already loading you into a trebuchet ("please hold still, we need to strap you in").

As they're making the final adjustments, they discover that their calibration document (your puzzle input) has been amended by a very young Elf who was apparently just excited to show off her art skills. Consequently, the Elves are having trouble reading the values on the document.

The newly-improved calibration document consists of lines of text; each line originally contained a specific calibration value that the Elves now need to recover. On each line, the calibration value can be found by combining the first digit and the last digit (in that order) to form a single two-digit number.

For example:

1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet

In this example, the calibration values of these four lines are 12, 38, 15, and 77. Adding these together produces 142.

Consider your entire calibration document. What is the sum of all of the calibration values?
*/
func SolvePart1(input string) int {
	sum := 0

	// First, split the input string into separate lines.
	lines := strings.Split(input, "\n")

	// Loop through each line to find the numbers.
	for _, line := range lines {
		var firstDigit, lastDigit = "", ""

		for _, char := range line {
			if unicode.IsDigit(char) {
				// Attempt to match the first character first.
				// If we haven't set it yet, set as both the first and last.
				if firstDigit == "" {
					firstDigit = string(char)
					lastDigit = string(char)
				} else {
					// Otherwise, we're just updating the last.
					lastDigit = string(char)
				}
			}
		}

		numberString := strings.Join([]string{firstDigit, lastDigit}, "")
		num, err := strconv.Atoi(numberString)
		if err != nil {
			panic(fmt.Sprintf("unable to resolve to number %s", numberString))
		}

		sum += num
	}

	return sum
}

/*
*
Your calculation isn't quite right. It looks like some of the digits are actually spelled out with letters: one, two, three, four, five, six, seven, eight, and nine also count as valid "digits".

Equipped with this new information, you now need to find the real first and last digit on each line. For example:

two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen

In this example, the calibration values are 29, 83, 13, 24, 42, 14, and 76. Adding these together produces 281.

What is the sum of all of the calibration values?
*/
func SolvePart2(input string) int {
	sum := 0

	// First, split the input string into separate lines.
	lines := strings.Split(input, "\n")

	// Loop through each line to find the numbers.
	for _, line := range lines {
		var firstDigit, firstIdx, lastDigit = "", 0, ""
		var writtenSequence = ""

		// The problem description didn't say this explicitly, nor does the test
		// data exemplify it, but we need to match the first digit greedily
		// from left to right, and the last digit greedily from right to left.
		// It's okay if there's an overlap, e.g., eighthree = 83

		// First, go forward.
		for idx, char := range line {
			var foundDigit = ""

			if unicode.IsDigit(char) {
				foundDigit = string(char)
			} else {
				// Append to the written sequence and look for a number spelled out.
				writtenSequence += string(char)
				if spelled, digit := extractDigit(writtenSequence); spelled != "" {
					foundDigit = digit
				}
			}

			// If we have arrived at a numerical character, set/update the first
			// and last characters.
			if foundDigit != "" {
				firstDigit = foundDigit
				firstIdx = idx

				// Reset the sequence.
				writtenSequence = ""

				break
			}
		}

		// Next, go backward.
		for idx := len(line) - 1; idx >= 0; idx-- {
			char := line[idx]
			var foundDigit = ""
			var spelledLength = 0

			if unicode.IsDigit(rune(char)) {
				foundDigit = string(char)
			} else {
				writtenSequence = string(char) + writtenSequence
				if spelled, digit := extractDigit(writtenSequence); spelled != "" {
					foundDigit = digit
					spelledLength = len(spelled)
				}
			}

			if foundDigit != "" && (idx+spelledLength) > firstIdx {
				lastDigit = foundDigit
				break
			}
		}

		numberString := strings.Join([]string{firstDigit, lastDigit}, "")
		num, err := strconv.Atoi(numberString)
		if err != nil {
			panic(fmt.Sprintf("unable to resolve to number %s", numberString))
		}

		fmt.Printf("line=%s, num=%d\n", line, num)
		sum += num
	}

	return sum
}

func extractDigit(sequence string) (string, string) {
	// If the sequence is less than any of the words, return early.
	if len(sequence) < 3 {
		return "", ""
	}

	if strings.Contains(sequence, "one") {
		return "one", "1"
	} else if strings.Contains(sequence, "two") {
		return "two", "2"
	} else if strings.Contains(sequence, "three") {
		return "three", "3"
	} else if strings.Contains(sequence, "four") {
		return "four", "4"
	} else if strings.Contains(sequence, "five") {
		return "five", "5"
	} else if strings.Contains(sequence, "six") {
		return "six", "6"
	} else if strings.Contains(sequence, "seven") {
		return "seven", "7"
	} else if strings.Contains(sequence, "eight") {
		return "eight", "8"
	} else if strings.Contains(sequence, "nine") {
		return "nine", "9"
	}

	return "", ""
}
