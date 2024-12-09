package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, _ := os.Open("./input.txt")
	fileConverted, _ := io.ReadAll(file)
	fileContent := string(fileConverted)

	regexClean := regexp.MustCompile(`don't\(\).*?do\(\)`)
	cleanText := regexClean.ReplaceAllString(fileContent, "")

	regex := regexp.MustCompile(`mul\(\s*(\d+)\s*,\s*(\d+)\s*\)`)
	matches := regex.FindAllStringSubmatch(cleanText, -1)
	totalSum := 0
	for index := 0; index < len(matches); index++ {
		firstParameter, _ := strconv.Atoi(matches[index][1])
		secondParameter, _ := strconv.Atoi(matches[index][2])
		multiplication := firstParameter * secondParameter
		totalSum += multiplication
	}
	fmt.Println(totalSum)
}
