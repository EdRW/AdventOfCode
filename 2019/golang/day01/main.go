package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func requiredFuel(mass int) int {
	fuelMass := mass/3 - 2
	if fuelMass <= 0 {
		return 0
	}
	return fuelMass + requiredFuel(fuelMass)
}

func main() {
	inputPath := "day01/input.txt"

	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	totalFuel := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputTxt := scanner.Text()
		moduleMass, err := strconv.Atoi(inputTxt)
		if err != nil {
			log.Fatal(err)
			return
		}

		totalFuel += requiredFuel(moduleMass)
	}

	fmt.Printf("total fuel: %d\n", totalFuel)
}
