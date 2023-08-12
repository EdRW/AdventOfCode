package main

import (
	"fmt"
	"strconv"

	"aoc/utils"
)

func requiredFuel(mass int) int {
	fuelMass := mass/3 - 2
	if fuelMass <= 0 {
		return 0
	}
	return fuelMass + requiredFuel(fuelMass)
}

func main() {
	scanner, close := utils.NewFileScanner("day01/input.txt")
	defer close()

	totalFuel := 0

	for scanner.Scan() {
		inputTxt := scanner.Text()
		moduleMass := utils.OrDie(strconv.Atoi(inputTxt))
		totalFuel += requiredFuel(moduleMass)
	}

	fmt.Printf("total fuel: %d\n", totalFuel)
}
