package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
}

func main() {
	fmt.Println("Advent of Code 2025 - Day 5")

	ranges, ingredientIDs := readInput("input.txt")
	ranges = mergeRanges(ranges)

	fmt.Println("Part 1:", part1(ranges, ingredientIDs))
	fmt.Println("Part 2:", part2(ranges))
}

func readInput(filename string) ([]Range, []int) {
	file, _ := os.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ranges := []Range{}
	ingredientIDs := []int{}
	readingRanges := true

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// blank line = switch to reading ingredient IDs
		if line == "" {
			readingRanges = false
			continue
		}

		if readingRanges {
			rangeParts := strings.Split(line, "-")
			if len(rangeParts) == 2 {
				start, _ := strconv.Atoi(rangeParts[0])
				end, _ := strconv.Atoi(rangeParts[1])
				ranges = append(ranges, Range{start: start, end: end})
			}
		} else {
			id, _ := strconv.Atoi(line)
			ingredientIDs = append(ingredientIDs, id)
		}
	}

	return ranges, ingredientIDs
}

func mergeRanges(ranges []Range) []Range {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].start < ranges[j].start
	})

	merged := []Range{ranges[0]}
	for i := 1; i < len(ranges); i++ {
		last := &merged[len(merged)-1]
		current := ranges[i]

		if current.start <= last.end+1 {
			if current.end > last.end {
				last.end = current.end
			}
		} else {
			merged = append(merged, current)
		}
	}

	return merged
}

func part1(ranges []Range, ingredientIDs []int) int {
	count := 0
	for _, id := range ingredientIDs {
		for _, r := range ranges {
			if id >= r.start && id <= r.end {
				count++
				break
			}
		}
	}

	return count
}

func part2(ranges []Range) int {
	freshIngredients := 0
	for _, r := range ranges {
		freshIngredients += r.end - r.start + 1
	}
	return freshIngredients
}
