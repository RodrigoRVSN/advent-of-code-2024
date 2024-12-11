package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func locationHasXmas(splittedByLine []string, y int, x int, dy int, dx int) bool {
	suffix := "MAS"
	for suffixIndex := range suffix {
		newY := y + (dy * (suffixIndex + 1))
		newX := x + (dx * (suffixIndex + 1))
		if newY < 0 || newX < 0 || newY >= len(splittedByLine)-1 || newX >= len(splittedByLine[0]) {
			return false
		}

		if splittedByLine[newY][newX] != suffix[suffixIndex] {
			return false
		}
	}
	return true
}

func main() {
	file, _ := os.Open("./input.txt")
	fileConverted, _ := io.ReadAll(file)
	fileContent := string(fileConverted)
	splittedByLine := strings.Split(fileContent, "\n")

	xmasCount := 0
	for y, line := range splittedByLine {
		for x, character := range line {
			if character == 'X' {
				for _, dy := range [3]int{-1, 0, 1} {
					for _, dx := range [3]int{-1, 0, 1} {
						if locationHasXmas(splittedByLine, y, x, dy, dx) {
							xmasCount++
						}
					}
				}
			}
		}
	}
	fmt.Println("Total Count:", xmasCount)
}
