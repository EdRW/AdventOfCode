package main

import (
	"aoc/machine"
	"aoc/utils"
	"fmt"
)

func main() {
	inputPath := utils.AOCInputFile(5)
	intCodes := machine.GetIntCodesFromFile(inputPath)

	// configure the computer
	options := machine.Options{
		IO: &machine.IO{
			StdIn:  machine.NewUserInputReader("Enter the System ID: "),
			StdOut: machine.NewUserOutputWriter("Output: "),
		},
	}
	computer := machine.NewComputer(options)

	// part 1
	fmt.Println("~~~~~ Part 1 ~~~~~")
	fmt.Println("HINT: Enter '1' for System ID")
	computer.Run(intCodes)
	fmt.Println("expected:", 13547311)

	fmt.Println()

	// part 2
	fmt.Println("~~~~~ Part 2 ~~~~~")
	fmt.Println("HINT: Enter '5' for System ID")
	computer.Run(intCodes)
	fmt.Println("expected:", 236453)
}
