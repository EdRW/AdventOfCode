package main

import (
	"aoc/machine"
	"aoc/utils"
	"fmt"
)

func getPermutations[T any](values []T) [][]int {
	// how to get every combination of phases,
	// without reusing any of the phase numbers?
	// 01234, 10234, 12034, 12304, 12340
	// 12340, 10234, 12034, ..., 12340

	// normally it might look like this if we could reuse
	// 00000, 10000, 20000, 30000, 40000
	// 01000,01000,01000,01000,01000,
	// 2,
	// but we can't because we're using the same phase
	return make([][]int, 0)
}

func part1() {
	NUM_PHASES := 5
	NUM_AMPS := 5

	phaseValues := make([]int, 5)
	for i := 0; i < NUM_PHASES; i++ {
		phaseValues[i] = i
	}

	phaseCombos := getPermutations(phaseValues)

	amps := make([]*machine.Computer, 5)
	for i := 0; i < NUM_AMPS; i++ {
		inPipe := machine.NewPipe(2)
		outPipe := machine.NewPipe(2)
		io := &machine.IO{
			StdIn:  inPipe,
			StdOut: outPipe,
		}
		amps[i] = machine.NewComputer(machine.Options{
			IO: io,
		})
	}

	maxOutput := 0
	for _, phaseCombo := range phaseCombos {
		input := 0
		for i, phase := range phaseCombo {
			amp := amps[i]
			amp.Input(phase)
			amp.Input(input)
			input = amp.Output()
		}
		maxOutput = utils.Max(maxOutput, input)
	}
}

func main() {
	fmt.Printf("Output: %d\n", 1)
	var userInput int
	utils.OrDie(fmt.Scanf("%d", &userInput))
	fmt.Printf("Output: %d\n", userInput)

}
