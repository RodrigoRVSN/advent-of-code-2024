package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func isSequenceSafe(isDecreasing bool, currentLevel int, previousLevel int) bool {
	if currentLevel-previousLevel == 0 {
		return false
	}
	if isDecreasing && currentLevel > previousLevel {
		return false
	}
	if !isDecreasing && currentLevel < previousLevel {
		return false
	}
	if isDecreasing {
		difference := previousLevel - currentLevel
		if difference > 3 {
			return false
		}
	}
	difference := currentLevel - previousLevel
	return difference <= 3
}

func main() {
	file, _ := os.Open("./input.txt")
	fileConverted, _ := io.ReadAll(file)
	fileContent := string(fileConverted)
	reports := strings.Split(fileContent, "\n")

	safeCount := 0
	for index := 0; index < len(reports)-1; index++ {
		levelsByReport := strings.Split(reports[index], " ")
		isSafe := true
		previousLevel, _ := strconv.Atoi(levelsByReport[0])
		secondLevel, _ := strconv.Atoi(levelsByReport[1])
		isDecreasing := secondLevel < previousLevel
		for levelIndex := 1; levelIndex < len(levelsByReport); levelIndex++ {
			currentLevel, _ := strconv.Atoi(levelsByReport[levelIndex])
			isSafe = isSequenceSafe(isDecreasing, currentLevel, previousLevel)
			previousLevel = currentLevel
			if !isSafe {
				break
			}
		}
		if isSafe {
			safeCount++
		}
	}

	fmt.Println("Safe count: ", safeCount)
}
