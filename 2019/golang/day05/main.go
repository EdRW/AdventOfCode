package main

import (
	"aoc/machine"
	"aoc/utils"
	"fmt"
)

func main() {
	inputPath := utils.AOCInputFile(5)
	intCodes := machine.GetIntCodesFromFile(inputPath)

	// part 1
	fmt.Println("~~~~~ Part 1 ~~~~~")
	fmt.Println("HINT: Enter '1' for System ID")
	computer := machine.NewComputer()
	computer.Run(intCodes)

	fmt.Println()

	// part 2
	fmt.Println("~~~~~ Part 2 ~~~~~")
	fmt.Println("HINT: Enter '5' for System ID")
	computer.Run(intCodes)
}
