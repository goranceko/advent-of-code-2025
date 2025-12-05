package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Advent of Code 2025 - Day 4\n")

	part1()
	part2()
}

func part1() {
	totalSum := 0

	file, _ := os.Open("input.txt")
	defer file.Close()

	var grid []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	rows := len(grid)
	cols := len(grid[0])

	dirs := [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] != '@' {
				continue
			}
			count := 0
			for _, d := range dirs {
				nr, nc := r+d[0], c+d[1]
				if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] == '@' {
					count++
				}
			}
			if count < 4 {
				totalSum++
			}
		}
	}

	fmt.Println("Part 1:", totalSum)
}

func part2() {
	fmt.Println("Part 2:", 0)
}
