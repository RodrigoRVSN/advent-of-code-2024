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

func getXmasCount(lines []string) int {
	xmasCount := 0
	for y, line := range lines {
		for x, character := range line {
			if character == 'X' {
				for _, dy := range [3]int{-1, 0, 1} {
					for _, dx := range [3]int{-1, 0, 1} {
						if locationHasXmas(lines, y, x, dy, dx) {
							xmasCount++
						}
					}
				}
			}
		}
	}
	return xmasCount
}

func locationHasCrossedMas(lines []string, y int, x int) bool {
	if y-1 < 0 || x-1 < 0 || y+1 >= len(lines)-1 || x+1 >= len(lines[0]) {
		return false
	}

	ulbr := string(lines[y-1][x-1]) + string(lines[y+1][x+1])
	if ulbr != "MS" && ulbr != "SM" {
		return false
	}

	urbl := string(lines[y-1][x+1]) + string(lines[y+1][x-1])
	if urbl != "MS" && urbl != "SM" {
		return false
	}
	return true
}

func getCrossedMasCount(lines []string) int {
	xmasCount := 0
	for y, line := range lines {
		for x, character := range line {
			if character == 'A' {
				if locationHasCrossedMas(lines, y, x) {
					xmasCount++
				}
			}
		}
	}
	return xmasCount
}

func main() {
	file, _ := os.Open("./input.txt")
	fileConverted, _ := io.ReadAll(file)
	fileContent := string(fileConverted)
	lines := strings.Split(fileContent, "\n")

	fmt.Println("Total Count of Xmas: ", getXmasCount(lines))
	fmt.Println("Total Count of Crossed Mas: ", getCrossedMasCount(lines))
}
