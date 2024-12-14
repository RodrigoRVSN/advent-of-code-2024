package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}

func sortSequence(values []string, rules []string) []string {
	for {
		swapped := false
		for _, rule := range rules {
			ruleSide := strings.Split(rule, "|")
			left := indexOf(ruleSide[0], values)
			right := indexOf(ruleSide[1], values)
			if left > right && right != -1 && left != -1 {
				values[left], values[right] = values[right], values[left]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return values
}

func main() {
	file, _ := os.Open("./input.txt")
	fileConverted, _ := io.ReadAll(file)
	fileContent := string(fileConverted)
	cleanContent := strings.TrimSpace(fileContent)
	sections := strings.Split(cleanContent, "\n\n")
	rules := strings.Split(sections[0], "\n")
	sequences := strings.Split(sections[1], "\n")
	totalCount := 0
	totalOrderedCount := 0

	for _, sequence := range sequences {
		isValidSequence := true
		previousSequences := []string{}
		values := strings.Split(sequence, ",")
		for valueIndex, value := range values {
			previousSequences = append(previousSequences, value)
			for _, rule := range rules {
				ruleSide := strings.Split(rule, "|")
				if ruleSide[1] == value {
					index := indexOf(ruleSide[0], values)
					if index > valueIndex {
						isValidSequence = false
					}
				}
			}
		}

		if !isValidSequence {
			sortedValues := sortSequence(values, rules)
			middleIndex := len(sortedValues) / 2
			valueAsNumber, _ := strconv.Atoi(string(sortedValues[middleIndex]))
			totalOrderedCount += valueAsNumber
		}

		if isValidSequence {
			middleIndex := len(values) / 2
			valueAsNumber, _ := strconv.Atoi(string(values[middleIndex]))
			totalCount += valueAsNumber
		}
	}

	fmt.Println("Total: ", totalCount)
	fmt.Println("Total after ordering it: ", totalOrderedCount)
}
