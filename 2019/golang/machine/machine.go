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
	// ctx                ExecutionContext
	instructionPointer int
	memory             Memory
	output             int
	io                 IO
	// envVars            utils.Set[string]
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

type Options struct {
	IO *IO
	// envVars *utils.Set[string]
}

func NewComputer(opts ...Options) *Computer {
	var options Options
	if len(opts) > 0 {
		options = opts[0]
	}

	if options.IO == nil {
		options.IO = &IO{
			StdIn:  StdIn,
			StdOut: StdOut,
		}
	}

	return &Computer{io: *options.IO}
}

func (c *Computer) advanceInstructionPointer(size int) {
	c.instructionPointer += size
}

func (c *Computer) process() bool {
	instruction := NewInstruction(
		NewExecutionContext(c.io.StdIn, c.io.StdOut),
		c.instructionPointer, &c.memory)

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

func (c *Computer) Input(value int) {
	c.io.StdIn.Write(value)
}

func (c *Computer) Output() int {
	return c.io.StdOut.Read()
}
