package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Advent of Code 2025 - Day 6")

	fmt.Println("Part 1:", part1())
	fmt.Println("\nPart 2:", part2())
}

func part1() int64 {
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	var sum int64
	operations := strings.Fields(lines[4])

	rows := make([][]string, 4)
	for i := 0; i < 4; i++ {
		rows[i] = strings.Fields(lines[i])
	}

	for col := 0; col < len(operations); col++ {
		num1 := num(rows[0][col])
		num2 := num(rows[1][col])
		num3 := num(rows[2][col])
		num4 := num(rows[3][col])

		switch operations[col] {
		case "+":
			sum += num1 + num2 + num3 + num4
		case "*":
			sum += num1 * num2 * num3 * num4
		}
	}

	return sum
}

func num(s string) int64 {
	var num int64
	fmt.Sscanf(s, "%d", &num)
	return num
}

func part2() int64 {
	return 0
}
