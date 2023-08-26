package machine

import (
	"log"
)

type Instruction struct {
	op        Operation
	paramVars []Variable
}

func NewInstruction(instructionPointer int, memory *Memory) Instruction {
	opCode, paramModes := parseOpCode(memory.deref(instructionPointer))

	op, ok := OpMap[opCode]
	if !ok {
		log.Fatalf("Unsupported operation: %d", int(opCode))
	}

	params := memory.slice(instructionPointer+1, instructionPointer+1+op.numParams)

	paramVars := NewVariables(memory, params, paramModes)

	return Instruction{
		op:        op,
		paramVars: paramVars,
	}
}

func (i *Instruction) Size() int {
	return i.op.numParams + 1
}

func (i *Instruction) Exec() bool {
	if i.op.opCode == Halt {
		return true
	}

	i.op.run(i.paramVars...)

	return false
}
