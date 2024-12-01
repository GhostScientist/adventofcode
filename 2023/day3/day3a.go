package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	sum, err := calculateSumOfPartNumbers("day3-input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Sum of part numbers:", sum)
}

func calculateSumOfPartNumbers(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	var schematic []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		schematic = append(schematic, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return sumPartNumbers(schematic), nil
}

func sumPartNumbers(schematic []string) int {
	sum := 0
	for y, line := range schematic {
		for x := 0; x < len(line); x++ {
			if unicode.IsDigit(rune(line[x])) {
				number, length := extractNumber(x, y, line)
				if isAdjacentToSymbol(x, y, schematic, length) {
					sum += number
					x += length - 1 // Skip the rest of this number
				}
			}
		}
	}
	return sum
}

func extractNumber(x, y int, line string) (int, int) {
	numberStr := ""
	length := 0
	for i := x; i < len(line) && unicode.IsDigit(rune(line[i])); i++ {
		numberStr += string(line[i])
		length++
	}
	number, _ := strconv.Atoi(numberStr)
	return number, length
}

func isAdjacentToSymbol(x, y int, schematic []string, length int) bool {
	for i := 0; i < length; i++ {
		if hasAdjacentSymbol(x+i, y, schematic) {
			return true
		}
	}
	return false
}

func hasAdjacentSymbol(x, y int, schematic []string) bool {
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			nx, ny := x+dx, y+dy
			if ny >= 0 && ny < len(schematic) && nx >= 0 && nx < len(schematic[ny]) {
				neighbor := rune(schematic[ny][nx])
				if neighbor != '.' && !unicode.IsDigit(neighbor) {
					return true
				}
			}
		}
	}
	return false
}
