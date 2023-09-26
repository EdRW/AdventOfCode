package machine

import "log"

type Mode int

const (
	Position Mode = iota
	Immediate
	Relative
)

type Variable struct {
	memory       *Memory
	valOrAddr    int
	mode         Mode
	relativeBase int
}

func (v *Variable) Get() int {
	switch v.mode {
	case Position:
		return v.memory.deref(v.valOrAddr)
	case Immediate:
		return v.valOrAddr
	case Relative:
		return v.memory.deref(v.relativeBase + v.valOrAddr)
	default:
		log.Fatalf("Unsupported parameter mode: %d\n", v.mode)
		return 0
	}
}

func (v *Variable) Set(value int) {
	switch v.mode {
	case Position:
		v.memory.assign(v.valOrAddr, value)
	case Immediate:
		log.Fatalln("Cannot set value of immediate variable")
	case Relative:
		v.memory.assign(v.relativeBase+v.valOrAddr, value)
	default:
		log.Fatalf("Unsupported parameter mode: %d\n", v.mode)
	}
}

func NewVariable(memory *Memory, param int, paramMode Mode, relativeBase int) Variable {
	return Variable{
		memory:       memory,
		valOrAddr:    param,
		mode:         paramMode,
		relativeBase: relativeBase,
	}
}

func NewVariables(
	memory *Memory,
	params []int,
	paramModes []int,
	relativeBase int,
) []Variable {
	vars := make([]Variable, len(params))

	for index, param := range params {
		vars[index] = NewVariable(
			memory,
			param,
			getParamMode(index, paramModes),
			relativeBase,
		)
	}

	return vars
}

func getParamMode(paramIndex int, paramModes []int) Mode {
	numParamModes := len(paramModes)
	if paramIndex >= numParamModes {
		return 0
	}
	return Mode(paramModes[paramIndex])
}
