package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	totalPower, err := calculateTotalPowerOfMinimumSets("day2-input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Total power of minimum sets:", totalPower)
}

func calculateTotalPowerOfMinimumSets(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalPower := 0
	for scanner.Scan() {
		line := scanner.Text()
		power := calculatePowerOfMinimumSet(line)
		fmt.Println("Power of minimum set:", power)
		totalPower += power
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return totalPower, nil
}

func calculatePowerOfMinimumSet(game string) int {
	minRed, minGreen, minBlue := 0, 0, 0
	// Corrected regular expression pattern
	pattern := regexp.MustCompile(`(\d+) (red|green|blue)`)
	matches := pattern.FindAllStringSubmatch(game, -1)

	for _, match := range matches {
		fmt.Println("Match:", match) // Debugging line
		count, _ := strconv.Atoi(match[1])
		switch match[2] {
		case "red":
			if count > minRed {
				minRed = count
			}
		case "green":
			if count > minGreen {
				minGreen = count
			}
		case "blue":
			if count > minBlue {
				minBlue = count
			}
		}
	}

	return minRed * minGreen * minBlue
}
