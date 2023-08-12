package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func orDie[T any](val T, err error) T {
	if err != nil {
		log.Fatal(err)
	}
	return val
}

// NewFileScanner opens a file and
// returns an associated *bufio.Scanner
// as well as a func to close the file.
//
//	scanner, close := NewFileScanner("day01/input.txt")
//	defer close()
func NewFileScanner(inputPath string) (scanner *bufio.Scanner, close func() error) {
	file := orDie(os.Open(inputPath))
	return bufio.NewScanner(file), file.Close
}

func requiredFuel(mass int) int {
	fuelMass := mass/3 - 2
	if fuelMass <= 0 {
		return 0
	}
	return fuelMass + requiredFuel(fuelMass)
}

func main() {
	scanner, close := NewFileScanner("day01/input.txt")
	defer close()

	totalFuel := 0

	for scanner.Scan() {
		inputTxt := scanner.Text()
		moduleMass := orDie(strconv.Atoi(inputTxt))
		totalFuel += requiredFuel(moduleMass)
	}

	fmt.Printf("total fuel: %d\n", totalFuel)
}
