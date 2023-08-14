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

var OpFunc = map[OpCode]func(int, int) int{
	Add:      add,
	Multiply: multiply,
}

func (m *Machine) runOp(opCode OpCode, pointer1 int, pointer2 int, outPointer int) {
	op, ok := OpFunc[opCode]
	if !ok {
		log.Fatalf("Unsupported command: %d", int(opCode))
	}
	param1, param2 := m.deref(m.pointer+1), m.deref(m.pointer+2)
	output := op(param1, param2)
	fmt.Printf("instruction: %d[%d](%d[%d]:%d[%d]) => %d[%d]\n",
		opCode, m.pointer, param1, m.pointer+1, param2, m.pointer+2, output, m.pointer+3)

	m.assignValue(m.pointer+3, output)
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

func (m *Machine) assignValue(pointer int, value int) {
	m.intCodes[m.intCodes[pointer]] = value
}

// advance instruction pointer
func (m *Machine) advanceIntCodePointer() {
	m.pointer += m.size
}

func (m *Machine) process() bool {
	opCode := OpCode(m.intCodes[m.pointer])
	if opCode == Halt {
		m.output = m.intCodes[0]
		fmt.Println("Halted!")
		return false
	}

	m.runOp(opCode, m.pointer+1, m.pointer+2, m.pointer+3)

	m.advanceIntCodePointer()

	return true
}

func (m *Machine) Run() {
	for m.process() {
	}
}

func getIntCodes() []int {

	input := utils.AocInputFile(2)
	scanner, close := utils.NewFileScanner(input)
	if !scanner.Scan() {
		log.Fatal()
	}
	close()

	inputTxt := scanner.Text()
	intCodeStrs := strings.Split(inputTxt, ",")
	return utils.ToInts(intCodeStrs)

}

func main() {
	intCodes := getIntCodes()

	// Restoring the gravity assist program inputs
	intCodes[1] = 12
	intCodes[2] = 2
	fmt.Printf("intCodes: %d\n", intCodes)

	instructionSize := 4
	machine := NewMachine(instructionSize, intCodes)
	machine.Run()

	fmt.Printf("result: %d\n", machine.output)
}
