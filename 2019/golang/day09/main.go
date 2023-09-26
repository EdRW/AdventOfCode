package main

import (
	"aoc/machine"
	"aoc/utils"
	"fmt"
)

func main() {
	inputPath := utils.AOCInputFile(9)
	intCodes := machine.GetIntCodesFromFile(inputPath)

	// configure the computer
	computer := machine.NewComputer()

	// part 1
	fmt.Println("~~~~~ Part 1 ~~~~~")
	fmt.Println("HINT: Enter '1' for the Input")
	computer.Run(intCodes)
	fmt.Println()
	fmt.Println("expected:", 3780860499)

	// fmt.Println()

	// part 2
	// fmt.Println("~~~~~ Part 2 ~~~~~")
}
