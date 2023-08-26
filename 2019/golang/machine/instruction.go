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

	params := memory.slice(instructionPointer+1, instructionPointer+1+op.size)

	paramVars := NewVariables(memory, params, paramModes)

	return Instruction{
		op:        op,
		paramVars: paramVars,
	}
}

func (i *Instruction) Size() int {
	return i.op.size + 1
}

func (i *Instruction) Exec() OpResult {
	return i.op.run(i.paramVars...)
}
