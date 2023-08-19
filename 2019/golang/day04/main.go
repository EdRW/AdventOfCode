package main

import (
	"aoc/utils"
	"fmt"
)

func isValidPassword(password []int) bool {
	prevDigit := password[0]

	doubleFound := false

	numRepeats := 1

	for _, digit := range password[1:] {
		if digit < prevDigit {
			return false
		}

		if digit == prevDigit {
			numRepeats++
		} else {
			if numRepeats == 2 {
				doubleFound = true
			}
			numRepeats = 1
		}
		prevDigit = digit
	}

	if numRepeats == 2 {
		return true
	}

	return doubleFound
}

func main() {
	// rules
	// has adjacent doubles
	// left to right never decreases,
	// only increase or stays same
	min, max := 245182, 790572

	numValidPasswords := 0
	for i := min; i <= max; i++ {
		if isValidPassword(utils.IntToInts(i)) {
			numValidPasswords++
		}
	}

	fmt.Printf("number of valid passwords: %d\n", numValidPasswords)
}
