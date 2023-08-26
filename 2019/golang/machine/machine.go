package machine

type Memory []int

func (m *Memory) deref(address int) int {
	return (*m)[address]
}
func (m *Memory) assign(address int, value int) {
	(*m)[address] = value
}

func (m *Memory) slice(start int, end int) Memory {
	return (*m)[start:end]
}

type Computer struct {
	instructionPointer int
	memory             Memory
	output             int
}

func (c *Computer) init(intCodes []int, firstInstructionAddress ...int) {
	if len(firstInstructionAddress) > 0 {
		c.instructionPointer = firstInstructionAddress[0]
	} else {
		c.instructionPointer = 0
	}

	intCodesCopy := make([]int, len(intCodes))
	copy(intCodesCopy, intCodes)
	c.memory = intCodesCopy
}

func NewComputer() *Computer {
	return &Computer{}
}

func (c *Computer) advanceInstructionPointer(size int) {
	c.instructionPointer += size
}

func (c *Computer) process() bool {
	instruction := NewInstruction(c.instructionPointer, &c.memory)
	result := instruction.Exec()
	if result.halt {
		c.output = c.memory[0]
		return false
	}

	if result.jump {
		c.instructionPointer = result.jumpToAddr
	} else {
		c.advanceInstructionPointer(instruction.Size())
	}

	return true
}

func (c *Computer) Run(intCodes []int, firstInstructionAddress ...int) int {
	c.init(intCodes, firstInstructionAddress...)
	for c.process() {
	}
	return c.output
}
