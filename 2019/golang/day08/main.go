package main

import (
	"aoc/utils"
	"fmt"
	"log"
	"strings"
)

func main() {
	const width = 25
	const height = 6
	const size = width * height

	input := utils.AOCInputFile(8)
	blob := readInputFile(input)

	numLayers := len(blob) / size

	layers := make([][]int, numLayers)
	for i := 0; i < numLayers; i++ {
		start := i * size
		end := (i + 1) * size
		layers[i] = blob[start:end]
	}

	minZerosDigitCounts := countDigits(layers[0])
	for i := 1; i < len(layers); i++ {
		layer := layers[i]
		digitCounts := countDigits(layer)
		if digitCounts[0] < minZerosDigitCounts[0] {
			minZerosDigitCounts = digitCounts
		}
	}

	output := minZerosDigitCounts[1] * minZerosDigitCounts[2]
	fmt.Println("Part 1 Output:", output)
}

func readInputFile(filePath string) []int {
	scanner, close := utils.NewFileScanner(filePath)
	if !scanner.Scan() {
		log.Fatal()
	}
	close()

	inputTxt := scanner.Text()
	ints := strings.Split(inputTxt, "")
	return utils.ToInts(ints)
}

func countDigits(layer []int) [3]int {
	digitCounts := [3]int{0, 0, 0}

	for _, pixel := range layer {
		digitCounts[pixel]++
	}

	return digitCounts
}
