package machine

import (
	"log"
)

type OpCode int

const (
	Add OpCode = iota + 1
	Multiply
	Halt OpCode = 99
)

type OpFunc func(int, int) int

func add(param1 int, param2 int) int {
	return param1 + param2
}

func multiply(param1 int, param2 int) int {
	return param1 * param2
}

var OpFuncMap = map[OpCode]OpFunc{
	Add:      add,
	Multiply: multiply,
}

type Machine struct {
	pointer  int
	size     int
	intCodes []int
	output   int
}

func (m *Machine) runOp(opCode OpCode, pParam1 int, pParam2 int, pResult int) {
	op, ok := OpFuncMap[opCode]
	if !ok {
		log.Fatalf("Unsupported command: %d", int(opCode))
	}

	param1, param2 := m.deref(pParam1), m.deref(pParam2)
	result := op(param1, param2)

	m.assignValue(pResult, result)
}

func (m *Machine) init(intCodes []int, pStart ...int) {
	pointer := 0
	if len(pStart) > 0 {
		pointer = pStart[0]
	}
	m.pointer = pointer

	intCodesCopy := make([]int, len(intCodes))
	copy(intCodesCopy, intCodes)
	m.intCodes = intCodesCopy
}

func NewMachine(size int) *Machine {
	return &Machine{
		size: size,
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
		return false
	}

	m.runOp(opCode, m.pointer+1, m.pointer+2, m.pointer+3)

	m.advanceIntCodePointer()

	return true
}

func (m *Machine) Run(intCodes []int, pStart ...int) int {
	m.init(intCodes, pStart...)
	for m.process() {
	}
	return m.output
}
