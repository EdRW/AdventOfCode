package main

import (
	"fmt"
	"log"
	"strings"

	"aoc/utils"
)

type OpCode int

const (
	Add OpCode = iota + 1
	Multiply
	Halt OpCode = 99
)

func add(param1 int, param2 int) int {
	return param1 + param2
}

func multiply(param1 int, param2 int) int {
	return param1 * param2
}

var Ops = map[OpCode]func(int, int) int{
	Add:      add,
	Multiply: multiply,
}

type Machine struct {
	pointer  int
	size     int
	intCodes []int
	output   int
}

func NewMachine(size int, intCodes []int) *Machine {
	return &Machine{
		pointer:  0,
		size:     size,
		intCodes: intCodes,
	}
}

func (m *Machine) deref(pointer int) int {
	return m.intCodes[m.intCodes[pointer]]
}

func (m *Machine) setValue(pointer int, value int) {
	m.intCodes[m.intCodes[pointer]] = value
}

func (m *Machine) process() bool {
	opCode := OpCode(m.intCodes[m.pointer])
	if opCode == Halt {
		m.output = m.intCodes[0]
		fmt.Println("Halted!")
		return false
	}

	op, ok := Ops[opCode]
	if !ok {
		log.Fatalf("Unsupported command: %d", int(opCode))
		return false
	}

	param1, param2 := m.deref(m.pointer+1), m.deref(m.pointer+2)
	output := op(param1, param2)

	m.setValue(m.pointer+3, output)

	// advance instruction pointer
	m.pointer += m.size

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

	instructionSize := 4
	machine := NewMachine(instructionSize, intCodes)

	for machine.process() {
	}

	fmt.Printf("result: %d\n", machine.output)
}
