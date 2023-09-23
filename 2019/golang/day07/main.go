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
	//   add the permutations to the list of permutations
	//   move the first element in the list to the end of the list
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

func getPhaseSet(numPhases int, initialPhase ...int) []int {
	offset := 0
	if len(initialPhase) != 0 {
		offset = initialPhase[0]
	}

	phaseValues := make([]int, numPhases)
	for i := 0; i < numPhases; i++ {
		phaseValues[i] = i + offset
	}
	return phaseValues
}

func allHalted(amps []*machine.Computer) bool {
	return utils.All(amps,
		func(amp *machine.Computer) bool {
			return amp.HasHalted()
		})
}

func part1(intCodes []int) int {
	NUM_PHASES := 5
	NUM_AMPS := 5

	phaseSet := getPhaseSet(NUM_PHASES)
	phaseCombos := getPermutations(phaseSet)

	expectedNumPerms := utils.Factorial(NUM_PHASES)
	if len(phaseCombos) != expectedNumPerms {
		log.Fatalf("Expected %d permutations, but got %d permutations",
			expectedNumPerms, len(phaseCombos))
	}

	// make the amplifiers
	amps := make([]*machine.Computer, NUM_AMPS)
	inputPipes := make([]machine.Pipe, NUM_AMPS)
	// the computers will use custom pipes for IO
	// rather than the default StdIn and StdOut
	firstInPipe := machine.NewPipe(2)
	thrusterPipe := machine.NewPipe(1)

	inPipe := firstInPipe
	for i := 0; i < NUM_AMPS; i++ {
		inputPipes[i] = inPipe

		// the final amp is a special case
		// since it outputs to thrusters
		var outPipe machine.Pipe
		if i == NUM_AMPS-1 {
			outPipe = thrusterPipe
		} else {
			outPipe = machine.NewPipe(2)
		}

		// set the input and output pipes
		options := machine.Options{
			Name: fmt.Sprint("Amp", i),
			IO: &machine.IO{
				StdIn:  inPipe,
				StdOut: outPipe,
			},
		}
		amps[i] = machine.NewComputer(options)

		inPipe = outPipe
	}
	// run the amplifier program on the amplifiers
	// try every combination of phases as inputs for the amps
	maxOutput := 0
	for _, phaseCombo := range phaseCombos {
		for i, phase := range phaseCombo {
			inPipe := inputPipes[i]
			inPipe.Write(phase)
		}

		// Set initial input for amp 1
		firstInPipe.Write(0)

		for _, amp := range amps {
			amp.Run(intCodes)
		}

		output := thrusterPipe.Read()
		maxOutput = utils.Max(maxOutput, output)
	}

	return maxOutput
}

func part2(intCodes []int) int {
	NUM_PHASES := 5
	NUM_AMPS := 5

	phaseSet := getPhaseSet(NUM_PHASES, 5)
	phaseCombos := getPermutations(phaseSet)

	expectedNumPerms := utils.Factorial(NUM_PHASES)
	if len(phaseCombos) != expectedNumPerms {
		log.Fatalf("Expected %d permutations, but got %d permutations",
			expectedNumPerms, len(phaseCombos))
	}

	// make the amplifiers
	amps := make([]*machine.Computer, NUM_AMPS)
	inputPipes := make([]machine.Pipe, NUM_AMPS)
	// the computers will use custom pipes for IO
	// rather than the default StdIn and StdOut
	firstInPipe := machine.NewPipe(2)
	thrusterPipe := machine.NewPipe(1)

	inPipe := firstInPipe
	for i := 0; i < NUM_AMPS; i++ {
		inputPipes[i] = inPipe

		// the final amp is a special case
		// since it outputs to amp 1 and thrusters
		var io *machine.IO
		if i == NUM_AMPS-1 {
			io = &machine.IO{
				StdIn:  inPipe,
				StdOut: machine.NewMultiWriter(firstInPipe, thrusterPipe),
			}
		} else {
			outPipe := machine.NewPipe(2)
			io = &machine.IO{
				StdIn:  inPipe,
				StdOut: outPipe,
			}
			inPipe = outPipe
		}

		// set machine options
		options := machine.Options{
			Name: fmt.Sprint("Amp", i),
			IO:   io,
		}
		amp := machine.NewComputer(options)
		amp.Init(intCodes)
		amps[i] = amp

	}

	// run the amplifier program on the amplifiers
	// try every combination of phases as inputs for the amps
	maxOutput := 0
	for _, phaseCombo := range phaseCombos {
		for i, phase := range phaseCombo {
			inPipe := inputPipes[i]
			inPipe.Write(phase)
		}

		// Set initial input for amp 1
		firstInPipe.Write(0)

		// returns true once all of the machines in a slice have halted

		for i := 0; !allHalted(amps); i++ {
			ampNum := i % NUM_AMPS
			amp := amps[ampNum]

			amp.RunWithInterrupts()

			if thrusterPipe.Len() > 0 {
				output := thrusterPipe.Read()
				maxOutput = utils.Max(maxOutput, output)
			}
		}

		// re-init the pipes and amps
		firstInPipe.Flush()
		for _, amp := range amps {
			amp.Init(intCodes)
		}
	}

	return maxOutput
}

func main() {
	inputPath := utils.AOCInputFile(7)
	intCodes := machine.GetIntCodesFromFile(inputPath)

	// part 1
	output := part1(intCodes)
	fmt.Println("Part 1 Output:", output)

	fmt.Println()

	// part 2
	output = part2(intCodes)
	fmt.Println("Part 2 Output:", output)
}
