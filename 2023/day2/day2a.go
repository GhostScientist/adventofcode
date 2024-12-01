package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	sum, err := calculateSumOfValidGameIDsFromFile("day2-input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Sum of IDs of valid games:", sum)
}

func calculateSumOfValidGameIDsFromFile(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		if isValidGame(line) {
			id, err := extractGameID(line)
			if err != nil {
				fmt.Println("Error extracting ID:", err)
				continue
			}
			sum += id
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return sum, nil
}

func extractGameID(game string) (int, error) {
	parts := strings.Split(game, ":")
	if len(parts) < 2 {
		return 0, fmt.Errorf("invalid game format")
	}
	idPart := strings.TrimSpace(parts[0])
	idStr := strings.Fields(idPart)
	if len(idStr) < 2 {
		return 0, fmt.Errorf("invalid ID format")
	}
	return strconv.Atoi(idStr[1])
}

func isValidGame(game string) bool {
	maxRed, maxGreen, maxBlue := 0, 0, 0
	pattern := regexp.MustCompile(`(\d+) (red|green|blue)`)
	matches := pattern.FindAllStringSubmatch(game, -1)

	for _, match := range matches {
		count, _ := strconv.Atoi(match[1])
		switch match[2] {
		case "red":
			if count > maxRed {
				maxRed = count
			}
		case "green":
			if count > maxGreen {
				maxGreen = count
			}
		case "blue":
			if count > maxBlue {
				maxBlue = count
			}
		}
	}

	return maxRed <= 12 && maxGreen <= 13 && maxBlue <= 14
}
