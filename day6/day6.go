package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Advent of Code 2025 - Day 6")

	// fmt.Println("Part 1:", part1())
	fmt.Println("\nPart 2:", part2())
}

func part1() int64 {
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	operations := strings.Fields(lines[4])

	rows := make([][]string, 4)
	for i := 0; i < 4; i++ {
		rows[i] = strings.Fields(lines[i])
	}

	var sum int64
	for col := 0; col < len(operations); col++ {
		result := num(rows[0][col])

		for row := 1; row < 4; row++ {
			val := num(rows[row][col])
			switch operations[col] {
			case "+":
				result += val
			case "*":
				result *= val
			}
		}
		sum += result
	}

	return sum
}

func num(s string) int64 {
	var num int64
	fmt.Sscanf(s, "%d", &num)
	return num
}

func part2() int64 {
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")

	var sum int64

	operations := lines[4]
	currentOp := operations[0]
	numbers := []int64{}
	value := ""
	for row := 0; row < 4; row++ {
		value += string(lines[row][0])
	}
	numbers = append(numbers, num(value))

	ops := []rune(operations)

	for i := 1; i < len(lines[0]); i++ {
		op := ops[i]
		switch op {
		case ' ':
			value := ""
			for row := 0; row < 4; row++ {
				value += string(lines[row][i])
			}
			if strings.TrimSpace(value) == "" {
				switch currentOp {
				case '+':
					result := int64(0)
					for _, n := range numbers {
						result += n
					}
					fmt.Println("result +:", result)
					sum += result
					fmt.Println("sum:", sum)
				case '*':
					result := int64(1)
					for _, n := range numbers {
						result *= n
					}
					fmt.Println("result *:", result)
					sum += result
					fmt.Println("sum:", sum)
				}
				numbers = nil
			} else {
				numbers = append(numbers, num(value))
			}

		case '+', '*':
			currentOp = uint8(op)
			value := ""
			for row := 0; row < 4; row++ {
				value += string(lines[row][i])
			}
			numbers = append(numbers, num(value))
		}
	}

	if currentOp == '+' {
		result := int64(0)
		for _, n := range numbers {
			result += n
		}
		sum += result
	} else {
		result := int64(1)
		for _, n := range numbers {
			result *= n
		}
		sum += result
	}

	return sum
}
