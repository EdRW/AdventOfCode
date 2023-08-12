package main

import (
	"fmt"
	"log"
	"strings"

	"aoc/utils"
)

type CommandName int

const (
	Add CommandName = iota + 1
	Multiply
	Halt CommandName = 99
)

type Machine struct {
	output int
}

func getIO(instruction []int) (int, int, int) {
	return instruction[1], instruction[2], instruction[3]
}

func (m *Machine) process(start int, intCodes []int) bool {
	optCode := CommandName(intCodes[start])
	if optCode == Halt {
		m.output = intCodes[0]
		fmt.Println("Halted!")
		return true
	}

	end := start + 4
	instruction := intCodes[start:end]
	in1, in2, out := getIO(instruction)
	fmt.Printf("instruction [%d:%d]: %v => ", start, end, instruction)

	if optCode == Add {
		sum := intCodes[in1] + intCodes[in2]
		fmt.Printf("(%d + %d = %d)\n", intCodes[in1], intCodes[in2], sum)
		intCodes[out] = sum
		return false
	}

	if optCode == Multiply {
		product := intCodes[in1] * intCodes[in2]
		fmt.Printf("(%d + %d = %d)\n", intCodes[in1], intCodes[in2], product)
		intCodes[out] = product
		return false
	}

	log.Fatalf("Unsupported command: %d", int(optCode))
	return true
}

func main() {
	input := utils.AocInputFile(2)
	scanner, close := utils.NewFileScanner(input)
	if !scanner.Scan() {
		log.Fatal()
	}
	close()

	inputTxt := scanner.Text()
	intCodeStrs := strings.Split(inputTxt, ",")
	intCodes := utils.ToInts(intCodeStrs)

	// Restoring the gravity assist program inputs
	intCodes[1] = 12
	intCodes[2] = 2
	fmt.Printf("intCodes: %d\n", intCodes)

	index := 0
	machine := &Machine{}

	for index < len(intCodes) {

		done := machine.process(index, intCodes)
		if done {
			break
		}
		index += 4
	}

	fmt.Printf("result: %d\n", machine.output)
}
