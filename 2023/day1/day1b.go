package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {

	file, err := os.Open("day1-input.txt")
	if err != nil {
		fmt.Println("Error reading file")
		return
	}
	defer file.Close()

	wordToDigit := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	var result int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		// Replace words with digits using the mapping
		for word, digit := range wordToDigit {
			line = strings.Replace(line, word, digit, -1)
			fmt.Println("Line after replacing", word, "with", digit, ":", line)
		}

		var firstDigit, lastDigit int
		foundFirstDigit := false

		for _, char := range line {
			if unicode.IsDigit(char) {
				digit, _ := strconv.Atoi(string(char))
				if !foundFirstDigit {
					firstDigit = digit
					foundFirstDigit = true
				}
				lastDigit = digit
			}
		}

		if foundFirstDigit {
			result += firstDigit*10 + lastDigit
		} else {
			// If no digits found in the line, handle accordingly
			// For example, you can choose to ignore or set to 0
			// result += 0
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Resulting sum:", result)
}
