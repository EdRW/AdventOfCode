package machine

import (
	"aoc/utils"
	"log"
	"strings"
)

type Memory []int

func (m *Memory) deref(address int) int {
	return (*m)[address]
}
func (m *Memory) assign(address int, value int) {
	(*m)[address] = value
}

type Machine struct {
	instructionPointer int
	memory             Memory
	output             int
}

func (m *Machine) init(intCodes []int, firstInstructionAddress ...int) {
	if len(firstInstructionAddress) > 0 {
		m.instructionPointer = firstInstructionAddress[0]
	} else {
		m.instructionPointer = 0
	}

	intCodesCopy := make([]int, len(intCodes))
	copy(intCodesCopy, intCodes)
	m.memory = intCodesCopy
}

func NewMachine() *Machine {
	return &Machine{}
}

func (m *Machine) advanceInstructionPointer(size int) {
	m.instructionPointer += size
}

func (m *Machine) process() bool {
	instruction := NewInstruction(m.instructionPointer, &m.memory)
	if instruction.Exec() {
		m.output = m.memory[0]
		return false
	}

	m.advanceInstructionPointer(instruction.Size())

	return true
}

func (m *Machine) Run(intCodes []int, firstInstructionAddress ...int) int {
	m.init(intCodes, firstInstructionAddress...)
	for m.process() {
	}
	return m.output
}

func GetIntCodesFromFile(filePath string) []int {
	scanner, close := utils.NewFileScanner(filePath)
	if !scanner.Scan() {
		log.Fatal()
	}
	close()

	inputTxt := scanner.Text()
	intCodeStrs := strings.Split(inputTxt, ",")
	return utils.ToInts(intCodeStrs)
}
