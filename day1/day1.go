package main

import (
	"bufio"
	"fmt"
	"os"
)

func openFileScanner() (*os.File, *bufio.Scanner) {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	return file, scanner
}

func part1() {
	dial := 50
	zerosCounter := 0

	file, scanner := openFileScanner()
	defer file.Close()

	for scanner.Scan() {
		line := scanner.Text()
		var direction rune
		var degrees int
		fmt.Sscanf(line, "%c%d", &direction, &degrees)

		if direction == 'L' {
			dial -= degrees
		} else {
			dial += degrees
		}

		// Ensure dial is always in range 0-99
		dial = ((dial % 100) + 100) % 100

		if dial == 0 {
			zerosCounter++
		}
	}
	fmt.Println("Part 1 - the password is:", zerosCounter) // 989
}

func part2() {
	dial := 50
	zerosCounter := 0

	file, scanner := openFileScanner()
	defer file.Close()
	for scanner.Scan() {
		line := scanner.Text()
		var direction rune
		var degrees int
		fmt.Sscanf(line, "%c%d", &direction, &degrees)

		for i := 0; i < degrees; i++ {
			if direction == 'L' {
				dial--
			} else {
				dial++
			}

			// Ensure dial is always in range 0-99
			dial = ((dial % 100) + 100) % 100

			if dial == 0 {
				zerosCounter++
			}
		}
	}

	fmt.Println("Part 2 - the password is:", zerosCounter) // 5941
}

func main() {
	part1()
	part2()
}
