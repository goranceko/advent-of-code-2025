package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Advent of Code 2025 - Day 3\n")
	content, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")

	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	sum := 0
	for _, line := range lines {
		numbers := []rune(line)
		d1 := int(slices.Max(numbers[:len(numbers)-1]) - '0')
		d2 := int(slices.Max(numbers[slices.Index(numbers, rune(d1)+'0')+1:]) - '0')
		sum += d1*10 + d2
	}
	fmt.Println("Part 1:", sum)
}

func part2(lines []string) {
	sum := 0
	for _, line := range lines {
		numbers := []rune(line)
		var selected string
		startIdx := 0

		for i := 0; i < 12; i++ {
			remainingNeeded := 12 - i - 1
			endIdx := len(numbers) - remainingNeeded
			maxRune := slices.Max(numbers[startIdx:endIdx])
			maxIdx := slices.Index(numbers[startIdx:endIdx], maxRune) + startIdx
			selected += string(maxRune)
			startIdx = maxIdx + 1
		}

		num, _ := strconv.Atoi(selected)
		sum += num
	}
	fmt.Println("Part 2:", sum)
}
