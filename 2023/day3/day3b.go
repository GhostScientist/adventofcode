package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	lines, err := readLines("day3-input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	totalSum := calculateTotalGearRatio(lines)
	fmt.Printf("Total gear ratio sum: %d\n", totalSum)
}

// Reads lines from the given file
func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// Calculates the total gear ratio
func calculateTotalGearRatio(lines []string) int {
	allNumbers := findAllNumbers(lines)
	totalSum := 0

	for lineIndex, line := range lines {
		starPositions := findCharPositions(line, '*')
		for _, starPos := range starPositions {
			foundNumbers := findNumbersAdjacentToGear(allNumbers, lineIndex, starPos, len(lines))
			if len(foundNumbers) == 2 {
				totalSum += foundNumbers[0].Number * foundNumbers[1].Number
			}
		}
	}

	return totalSum
}

// PosNumber holds the details of a found number
type PosNumber struct {
	StartPos  int
	EndPos    int
	LineIndex int
	Number    int
}

// Finds all numbers in the given lines
func findAllNumbers(lines []string) []*PosNumber {
	var allNumbers []*PosNumber
	re := regexp.MustCompile(`\d+`)

	for lineIndex, line := range lines {
		matches := re.FindAllStringIndex(line, -1)
		for _, match := range matches {
			number, _ := strconv.Atoi(line[match[0]:match[1]])
			allNumbers = append(allNumbers, &PosNumber{
				StartPos:  match[0],
				EndPos:    match[1],
				LineIndex: lineIndex,
				Number:    number,
			})
		}
	}

	return allNumbers
}

// Finds numbers adjacent to a gear
func findNumbersAdjacentToGear(allNumbers []*PosNumber, lineIndex, gearPos, totalLines int) []*PosNumber {
	var foundNumbers []*PosNumber
	for _, number := range allNumbers {
		if numberTouchesGear(number, lineIndex, gearPos, totalLines) {
			foundNumbers = append(foundNumbers, number)
		}
	}
	return foundNumbers
}

// Checks if a number is adjacent to a gear
func numberTouchesGear(number *PosNumber, gearLineIndex, gearPos, totalLines int) bool {
	for loopLine := gearLineIndex - 1; loopLine <= gearLineIndex+1; loopLine++ {
		if loopLine < 0 || loopLine >= totalLines || number.LineIndex != loopLine {
			continue
		}
		if number.StartPos <= gearPos && number.EndPos > gearPos {
			return true
		}
	}
	return false
}

// Finds positions of a specific character in a string
func findCharPositions(str string, targetChar rune) []int {
	positions := []int{}
	for i, char := range str {
		if char == targetChar {
			positions = append(positions, i)
		}
	}
	return positions
}
