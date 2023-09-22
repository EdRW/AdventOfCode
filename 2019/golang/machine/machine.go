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

type State int

const (
	Halted State = iota
	Running
	Paused
)

type Computer struct {
	// ctx                ExecutionContext
	name               string
	instructionPointer int
	memory             Memory
	output             int
	state              State
	io                 IO
	// envVars            utils.Set[string]
}

func (c *Computer) Init(intCodes []int, firstInstructionAddress ...int) {
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
	Name string
	IO   *IO
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
	if options.Name == "" {
		options.Name = "computer"
	}

	return &Computer{name: options.Name, io: *options.IO}
}

func (c *Computer) advanceInstructionPointer(size int) {
	c.instructionPointer += size
}

func (c *Computer) process() {
	instruction := NewInstruction(
		NewExecutionContext(c.io.StdIn, c.io.StdOut),
		c.instructionPointer, &c.memory)

	result := instruction.Exec()
	if result.halt {
		c.output = c.memory[0]
		c.state = Halted
		return
	} else if result.interrupt {
		c.state = Paused
	}

	if result.jump {
		c.instructionPointer = result.jumpToAddr
	} else {
		c.advanceInstructionPointer(instruction.Size())
	}
}

// Runs until the program halts, ignoring any interrupts
func (c *Computer) Run(intCodes []int, firstInstructionAddress ...int) int {
	c.Init(intCodes, firstInstructionAddress...)
	c.state = Running
	for c.state != Halted {
		c.process()
	}
	return c.output
}

// Runs until the program halts or an interrupt occurs
func (c *Computer) RunWithInterrupts() {
	c.state = Running
	for c.state == Running {
		c.process()
	}
}

func (c *Computer) HasHalted() bool {
	return c.state == Halted
}

func (c *Computer) Input(value int) {
	c.io.StdIn.Write(value)
}

func (c *Computer) Output() int {
	return c.io.StdOut.Read()
}
