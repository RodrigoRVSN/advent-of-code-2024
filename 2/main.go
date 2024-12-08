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

func remove(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}

func areLevelsSafe(initialLevels []string, currentLevels []string, indexTry int) bool {
	isSafe := true
	previousLevel, _ := strconv.Atoi(currentLevels[0])
	secondLevel, _ := strconv.Atoi(currentLevels[1])
	isDecreasing := secondLevel < previousLevel

	for levelIndex := 1; levelIndex < len(currentLevels); levelIndex++ {
		currentLevel, _ := strconv.Atoi(currentLevels[levelIndex])
		isSafe = isSequenceSafe(isDecreasing, currentLevel, previousLevel)
		previousLevel = currentLevel
		if indexTry < len(initialLevels) && !isSafe {
			newLevels := make([]string, len(initialLevels))
			copy(newLevels, initialLevels)
			newLevels = remove(newLevels, indexTry)
			return areLevelsSafe(initialLevels, newLevels, indexTry+1)
		}
		if !isSafe {
			break
		}
	}
	return isSafe
}

func main() {
	file, _ := os.Open("./input.txt")
	fileConverted, _ := io.ReadAll(file)
	fileContent := string(fileConverted)
	reports := strings.Split(fileContent, "\n")

	safeCount := 0
	for index := 0; index < len(reports)-1; index++ {
		levelsByReport := strings.Split(reports[index], " ")
		isSafe := areLevelsSafe(levelsByReport, levelsByReport, 0)
		if isSafe {
			safeCount++
		}
	}

	fmt.Println("Safe count: ", safeCount)
}
