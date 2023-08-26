package main

import (
	"aoc/machine"
	"aoc/utils"
	"fmt"
	"log"
)

func getIntCodes() []int {
	input := utils.AocInputFile(2)
	return machine.GetIntCodesFromFile(input)
}

func initIntCodes(intCodes []int, noun int, verb int) {
	intCodes[1] = noun
	intCodes[2] = verb
}

func bruteForceInputs(computer *machine.Machine, intCodes []int, goal int) (int, int) {
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			// Restoring the gravity assist program inputs
			initIntCodes(intCodes, i, j)

			output := computer.Run(intCodes)
			if output == goal {
				return i, j
			}
		}
	}
	log.Fatalf("unable to achieve goal %d\n", goal)
	return -1, -1
}

func main() {
	intCodes := getIntCodes()

	computer := machine.NewMachine()

	// part 1
	initIntCodes(intCodes, 12, 2)
	p1Output := computer.Run(intCodes)
	fmt.Printf("part 1 result: %d\n", p1Output)

	// part 2
	goal := 19690720
	noun, verb := bruteForceInputs(computer, intCodes, goal)

	// verify noun and verb
	initIntCodes(intCodes, noun, verb)
	if computer.Run(intCodes) != goal {
		log.Fatalf("Could not verify that noun: %d, verb: %d produce goal: %d\n", noun, verb, goal)
	}

	p2Output := 100*noun + verb

	fmt.Printf("part 2 result: %d\n", p2Output)
}
