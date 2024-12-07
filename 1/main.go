package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("./input.txt")
	fileConverted, _ := io.ReadAll(file)
	fileContent := string(fileConverted)
	splittedByLine := strings.Split(fileContent, "\n")
	left, right := []int{}, []int{}
	for index := 0; index < len(splittedByLine)-1; index++ {
		leftRight := strings.Split(splittedByLine[index], "   ")
		leftSide, _ := strconv.Atoi(leftRight[0])
		rightSide, _ := strconv.Atoi(leftRight[1])
		left = append(left, leftSide)
		right = append(right, rightSide)
	}
	sort.Ints(left)
	sort.Ints(right)
	count := 0.0
	for index := 0; index < len(splittedByLine)-1; index++ {
		count += math.Abs(float64(right[index]) - float64(left[index]))
	}
	fmt.Printf("%f", count)
}
