package day2

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const totalRed, totalGreen, totalBlue = 12, 13, 14

func SolvePart1(input string) int {
	var sum = 0
	// First, split out the total number of games.
	games := strings.Split(input, "\n")

GAMELOOP:
	for _, game := range games {
		// Split the line by the colon to get the games vs. the game number
		draws := strings.Split(strings.SplitAfter(game, ":")[1], ";")

		for _, draw := range draws {
			greenParams := getParams(`(?P<Number>\d+) green`, draw)
			redParams := getParams(`(?P<Number>\d+) red`, draw)
			blueParams := getParams(`(?P<Number>\d+) blue`, draw)

			if numString, ok := greenParams["Number"]; ok {
				num, err := strconv.Atoi(numString)
				if err != nil {
					panic(fmt.Sprintf("Unable to parse number: %s", numString))
				}

				if num > totalGreen {
					continue GAMELOOP
				}
			}
			if numString, ok := redParams["Number"]; ok {
				num, err := strconv.Atoi(numString)
				if err != nil {
					panic(fmt.Sprintf("Unable to parse number: %s", numString))
				}

				if num > totalRed {
					continue GAMELOOP
				}
			}
			if numString, ok := blueParams["Number"]; ok {
				num, err := strconv.Atoi(numString)
				if err != nil {
					panic(fmt.Sprintf("Unable to parse number: %s", numString))
				}

				if num > totalBlue {
					continue GAMELOOP
				}
			}
		}

		// If we made it this far, the game is legit.
		gameParams := getParams(`Game (?P<Game>\d+)`, game)
		gameNum, err := strconv.Atoi(gameParams["Game"])
		if err != nil {
			panic(fmt.Sprintf("Unable to parse game number: %s", gameParams["Game"]))
		}

		sum += gameNum
	}
	return sum
}

func SolvePart2(input string) int {
	var sum = 0
	// First, split out the total number of games.
	games := strings.Split(input, "\n")

	for _, game := range games {
		var maxGreen, maxRed, maxBlue = 0, 0, 0

		// Split the line by the colon to get the games vs. the game number
		draws := strings.Split(strings.SplitAfter(game, ":")[1], ";")

		for _, draw := range draws {
			greenParams := getParams(`(?P<Number>\d+) green`, draw)
			redParams := getParams(`(?P<Number>\d+) red`, draw)
			blueParams := getParams(`(?P<Number>\d+) blue`, draw)

			if numString, ok := greenParams["Number"]; ok {
				num, err := strconv.Atoi(numString)
				if err != nil {
					panic(fmt.Sprintf("Unable to parse number: %s", numString))
				}

				if num > maxGreen {
					maxGreen = num
				}
			}
			if numString, ok := redParams["Number"]; ok {
				num, err := strconv.Atoi(numString)
				if err != nil {
					panic(fmt.Sprintf("Unable to parse number: %s", numString))
				}

				if num > maxRed {
					maxRed = num
				}
			}
			if numString, ok := blueParams["Number"]; ok {
				num, err := strconv.Atoi(numString)
				if err != nil {
					panic(fmt.Sprintf("Unable to parse number: %s", numString))
				}

				if num > maxBlue {
					maxBlue = num
				}
			}
		}

		sum += maxGreen * maxRed * maxBlue
	}
	return sum
}

func getParams(regEx, url string) (paramsMap map[string]string) {

	var compRegEx = regexp.MustCompile(regEx)
	match := compRegEx.FindStringSubmatch(url)

	paramsMap = make(map[string]string)
	for i, name := range compRegEx.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
		}
	}
	return paramsMap
}
