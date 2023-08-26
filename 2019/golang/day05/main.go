package main

import (
	"aoc/machine"
	"aoc/utils"
	"fmt"
)

// TODO Day 5 machine updates
// [x] add new opcodes
// [x] parse opcode intCode to pull out opcode and  parameter modes
// [x] add support for respecting parameter modes
// [x] create an new type called Instruction that holds the opcode, parameter modes, and size
// [x] Instruction should should follow the command pattern
// [x] update the instruction pointer based on size of instruction
// [x] pass systemId as the first input
// [x] outputs from the program will represent the distance from the correct answer for each test
// [x] 0 means the test was OK
// [x] the final output before halting is the diagnostic code
// [x] all outputs before that should be a 0
// [x] the diagnostic code is the answer to this problem

func main() {
	inputPath := utils.AocInputFile(5)
	intCodes := machine.GetIntCodesFromFile(inputPath)

	// part 1
	fmt.Println("~~~~~ Part 1 ~~~~~")
	// input for System ID should be 1
	computer := machine.NewComputer()
	computer.Run(intCodes)
}
