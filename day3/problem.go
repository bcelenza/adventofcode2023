package day3

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

/*
*
You and the Elf eventually reach a gondola lift station; he says the gondola lift will take you up to the water source, but this is as far as he can bring you. You go inside.

It doesn't take long to find the gondolas, but there seems to be a problem: they're not moving.

"Aaah!"

You turn around to see a slightly-greasy Elf with a wrench and a look of surprise. "Sorry, I wasn't expecting anyone! The gondola lift isn't working right now; it'll still be a while before I can fix it." You offer to help.

The engineer explains that an engine part seems to be missing from the engine, but nobody can figure out which one. If you can add up all the part numbers in the engine schematic, it should be easy to work out which part is missing.

The engine schematic (your puzzle input) consists of a visual representation of the engine. There are lots of numbers and symbols you don't really understand, but apparently any number adjacent to a symbol, even diagonally, is a "part number" and should be included in your sum. (Periods (.) do not count as a symbol.)

Here is an example engine schematic:

467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..

In this schematic, two numbers are not part numbers because they are not adjacent to a symbol: 114 (top right) and 58 (middle right). Every other number is adjacent to a symbol and so is a part number; their sum is 4361.

Of course, the actual engine schematic is much larger. What is the sum of all of the part numbers in the engine schematic?
*/
type NumberLocation struct {
	x      int
	y      int
	length int
}

type GearLocation struct {
	x int
	y int
}

func Solve(input string, onlyGearRatios bool) int {
	var sum = 0

	// We'll need a place to store the locations of numbers.
	numberLocations := make([]NumberLocation, 0)

	// And a place to map gear locations with adjacent numbers for part 2.
	gearLocations := make(map[GearLocation][]int)

	// And a grid representing the locations of each character.
	grid := make([][]rune, 0)

	// Start by building the grid and extracting the number locations.
	// If we're counting gears for part 2, capture their locations too.
	lines := strings.Split(input, "\n")
	for y, line := range lines {
		gridLine := make([]rune, 0)

		var onNumber = false
		var numX = 0
		lineRunes := []rune(strings.TrimSpace(line))
		for x, char := range lineRunes {
			// Add the byte to the grid.
			gridLine = append(gridLine, char)

			// If this is the start of a number, record the index.
			if unicode.IsDigit(char) {
				if x == 0 || !unicode.IsDigit(lineRunes[x-1]) {
					onNumber = true
					numX = x
				}

				// If this is the end of the line, capture the number's location.
				if x == len(lineRunes)-1 {
					onNumber = false
					numberLocations = append(numberLocations, NumberLocation{
						x:      numX,
						y:      y,
						length: x - numX + 1,
					})
				}
			} else {
				// If char is a gear, capture its location
				if char == 42 {
					location := GearLocation{
						x: x,
						y: y,
					}
					gearLocations[location] = make([]int, 0)
				}

				// If we were on a number, we no longer are and can capture its
				// location.
				if onNumber {
					onNumber = false

					numberLocations = append(numberLocations, NumberLocation{
						x:      numX,
						y:      y,
						length: x - numX,
					})
				}
			}
		}

		grid = append(grid, gridLine)
	}

	// Okay, we have a grid and the locations of the numbers. Now we need to
	// see if any special characters are adjacent to the located numbers.
	for _, location := range numberLocations {
		// Extract the number in case we need it.
		num := ExtractNumber(lines, location)

		startX := location.x - 1
		if startX < 0 {
			startX = 0
		}

		// Search previous line if applicable.
		if location.y > 0 {
			gridLine := grid[location.y-1]
			endX := location.x + location.length + 1
			if endX > len(gridLine) {
				endX = len(gridLine)
			}

			gridSection := gridLine[startX:endX]

			if onlyGearRatios {
				gears := ExtractGearLocations(gridSection, startX, location.y-1)
				for _, gear := range gears {
					gearLocations[gear] = append(gearLocations[gear], num)
				}
			} else {
				if ContainsSpecialCharacter(gridLine[startX:endX]) {
					sum += num
					continue
				}
			}
		}

		// Search before and after in the current line.
		if location.x > 0 {
			gridSection := grid[location.y][location.x-1 : location.x]
			if onlyGearRatios {
				gears := ExtractGearLocations(gridSection, location.x-1, location.y)
				for _, gear := range gears {
					gearLocations[gear] = append(gearLocations[gear], num)
				}
			} else {
				if ContainsSpecialCharacter(gridSection) {
					sum += num
					continue
				}
			}
		}
		if location.x+location.length < len(lines[location.y])-1 {
			gridSection := grid[location.y][location.x+location.length : location.x+location.length+1]

			if onlyGearRatios {
				gears := ExtractGearLocations(gridSection, location.x+location.length, location.y)
				for _, gear := range gears {
					gearLocations[gear] = append(gearLocations[gear], num)
				}

			} else {
				if ContainsSpecialCharacter(gridSection) {
					sum += num
					continue
				}
			}
		}

		// Search next line if applicable.
		if location.y < len(grid)-1 {
			gridLine := grid[location.y+1]
			endX := location.x + location.length + 1
			if endX > len(gridLine) {
				endX = len(gridLine)
			}

			gridSection := gridLine[startX:endX]
			if onlyGearRatios {
				gears := ExtractGearLocations(gridSection, startX, location.y+1)
				for _, gear := range gears {
					gearLocations[gear] = append(gearLocations[gear], num)
				}
			} else {
				if ContainsSpecialCharacter(gridSection) {
					sum += num
					continue
				}
			}

		}

	}

	if onlyGearRatios {
		for _, nums := range gearLocations {
			if len(nums) == 2 {
				sum += nums[0] * nums[1]
			}
		}
	}

	return sum
}

func ContainsSpecialCharacter(slice []rune) bool {
	for _, r := range slice {
		if !unicode.IsDigit(r) && r != 46 {
			return true
		}
	}

	return false
}

func ExtractNumber(lines []string, location NumberLocation) int {
	line := lines[location.y]
	numString := line[location.x : location.x+location.length]
	num, err := strconv.Atoi(numString)
	if err != nil {
		panic(fmt.Sprintf("Unable to parse number: %s", numString))
	}
	return num
}

func ExtractGearLocations(slice []rune, startX int, y int) []GearLocation {
	locations := make([]GearLocation, 0)

	for x, char := range slice {
		if char == 42 {
			locations = append(locations, GearLocation{
				x: x + startX,
				y: y,
			})
		}
	}

	return locations
}
