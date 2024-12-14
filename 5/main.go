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

func main() {
	file, _ := os.Open("./input.txt")
	fileConverted, _ := io.ReadAll(file)
	fileContent := string(fileConverted)
	cleanContent := strings.TrimSpace(fileContent)
	sections := strings.Split(cleanContent, "\n\n")
	rules := strings.Split(sections[0], "\n")
	sequences := strings.Split(sections[1], "\n")
	totalCount := 0
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
						break
					}
				}
			}
		}
		if isValidSequence {
			middleIndex := len(values) / 2
			valueAsNumber, _ := strconv.Atoi(string(values[middleIndex]))
			totalCount += valueAsNumber
		}
	}
	fmt.Println(totalCount)
}
