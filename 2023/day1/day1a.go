package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {

	// Open the file for reading
	file, err := os.Open("day1-input.txt")
	if err != nil {
		fmt.Println("Error reading file")
		return
	}
	defer file.Close()

	// Scan each line of the file
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {

		// Get the string from the current line
		inputString := scanner.Text()

		// Find the first digit in the string
		var result int
		foundFirstDigit := false
		lastDigit := -1

		// Loop over all the characters in the input string
		for _, char := range inputString {
			// Check if the character is a digit
			if unicode.IsDigit(char) {
				// Convert the character to an integer
				digit := int(char - '0')
				// If this is the first digit, set the result
				if !foundFirstDigit {
					result = digit
					foundFirstDigit = true
				}
				// Set the last digit
				lastDigit = digit
			}
		}

		// If we found the first digit, add the last digit to the result
		if foundFirstDigit {
			result = result*10 + lastDigit
		}

		// Add the result to the sum
		sum += result

		fmt.Println("Resulting number:", result)
	}

	fmt.Println("Sum:", sum)
}
