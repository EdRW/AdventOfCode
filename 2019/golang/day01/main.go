package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func requiredFuel(moduleMass int) int {
	return moduleMass/3 - 2
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

		moduleFuel := requiredFuel(moduleMass)
		totalFuel += moduleFuel
	}

	fmt.Println(totalFuel)
}
