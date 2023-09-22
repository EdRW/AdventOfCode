package main

import (
	"aoc/machine"
	"aoc/utils"
	"fmt"
	"log"
)

func rotate[T any](values []T, times int) []T {
	front := values[:times]
	rest := values[times:]
	return append(rest, front...)
}

func getPermutations[T any](set []T) [][]T {
	// how to get every combination of phases,
	// without reusing any of the phase numbers?
	//
	// if the # elements in the list 1, return a list of list containing that element
	// make an empty list of permutations
	// for m of num elements in the list
	//   hold the first element of the list
	//   pass the rest of the list to recursive fn which gives the list of permutations of the sub-list
	//   prepend the first element to each sub-list permutation
	//   add the combined permutations to the list of permutations
	//   move the first m elements in the list to the end of the list
	// return the list of permutations
	//
	// 0123, 0132
	//       0231, 0213,
	//       0312, 0321,
	// 1230, 1203,
	//       1302, 1320,
	//       1023, 1032
	// 2301, 2310,
	// ...

	size := len(set)
	permutations := make([][]T, 0)

	if size == 0 {
		return permutations
	} else if size == 1 {
		permutations = append(permutations, set)
		return permutations
	}

	for i := 0; i < size; i++ {
		first := set[0]
		subset := utils.Copy(set[1:])

		subsetPermutations := getPermutations(subset)
		for _, subsetPermutation := range subsetPermutations {
			// prepend the first element to each subset permutation
			permutation := utils.Prepend(subsetPermutation, first)
			permutations = append(permutations, permutation)
		}

		// move the first element in the list to the end of the list
		set = rotate(set, 1)
	}

	return permutations
}

func part1(intCodes []int) int {
	NUM_PHASES := 5
	NUM_AMPS := 5

	phaseValues := make([]int, NUM_PHASES)
	for i := 0; i < NUM_PHASES; i++ {
		phaseValues[i] = i
	}
	phaseCombos := getPermutations(phaseValues)

	expectedNumPerms := utils.Factorial(NUM_PHASES)
	if len(phaseCombos) != expectedNumPerms {
		log.Fatalf("Expected %d permutations, but got %d permutations",
			expectedNumPerms, len(phaseCombos))
	}

	// make the amplifiers
	amps := make([]*machine.Computer, NUM_AMPS)
	for i := 0; i < NUM_AMPS; i++ {
		// the computers will use custom pipes for IO
		// rather than the default StdIn and StdOut
		inPipe := machine.NewPipe(2)
		outPipe := machine.NewPipe(2)

		// set the input and output pipes
		options := machine.Options{
			IO: &machine.IO{
				StdIn:  inPipe,
				StdOut: outPipe,
			},
		}
		amps[i] = machine.NewComputer(options)
	}

	// run the amplifier program on the amplifiers
	// try every combination of phases as inputs for the amps
	maxOutput := 0
	for _, phaseCombo := range phaseCombos {
		input := 0
		for i, phase := range phaseCombo {
			amp := amps[i]
			amp.Input(phase)
			amp.Input(input)
			amp.Run(intCodes)
			input = amp.Output()
		}
		maxOutput = utils.Max(maxOutput, input)
	}

	return maxOutput
}

func main() {

	inputPath := utils.AOCInputFile(7)
	intCodes := machine.GetIntCodesFromFile(inputPath)

	// part 1
	output := part1(intCodes)
	fmt.Println("Part 1 Output:", output)
}
