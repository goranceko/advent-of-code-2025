package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isInvalidIDPart1(id string) bool {
	if len(id)%2 != 0 {
		return false
	}
	mid := len(id) / 2
	return id[:mid] == id[mid:]
}

func isInvalidIDPart2(id string) bool {
	// Try pattern lengths from 1 to length/2
	for patternLen := 1; patternLen <= len(id)/2; patternLen++ {
		// is string length divisible by pattern length (eg. 9 % 5)
		if len(id)%patternLen != 0 {
			continue
		}

		pattern := id[:patternLen]
		repetitions := len(id) / patternLen
		if strings.Repeat(pattern, repetitions) == id {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println("Day 2 - Advent of Code 2025\n")

	data, _ := os.ReadFile("input.txt")
	input := strings.TrimSpace(string(data))
	ranges := strings.Split(input, ",")

	sum1 := 0
	sum2 := 0

	for _, rangePair := range ranges {
		parts := strings.Split(rangePair, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])

		for num := start; num <= end; num++ {
			id := strconv.Itoa(num)
			if isInvalidIDPart1(id) {
				sum1 += num
			}
			if isInvalidIDPart2(id) {
				sum2 += num
			}
		}
	}

	fmt.Printf("Part 1: Sum of invalid IDs: %d\n", sum1)
	fmt.Printf("Part 2: Sum of invalid IDs: %d\n", sum2)
}
