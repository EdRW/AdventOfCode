package machine

import "log"

type Mode int

const (
	Position Mode = iota
	Immediate
)

type Variable struct {
	memory    *Memory
	valOrAddr int
	mode      Mode
}

func (v *Variable) Get() int {
	if v.mode == Immediate {
		return v.valOrAddr
	}
	return v.memory.deref(v.valOrAddr)
}

func (v *Variable) Set(value int) {
	if v.mode == Immediate {
		log.Fatalf("Cannot set value of immediate variable")
	}
	v.memory.assign(v.valOrAddr, value)
}

func NewVariable(memory *Memory, param int, paramMode Mode) Variable {
	return Variable{
		memory:    memory,
		valOrAddr: param,
		mode:      paramMode,
	}
}

func NewVariables(memory *Memory, params []int, paramModes []int) []Variable {
	vars := make([]Variable, len(params))

	for index, param := range params {
		vars[index] = NewVariable(memory, param, getParamMode(index, paramModes))
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
