package machine

import (
	"aoc/utils"
)

type OpCode int

const (
	Add OpCode = iota + 1
	Multiply
	ReadStdIn
	WriteStdOut
	JmpIfTrue
	JmpIfFalse
	LessThan
	Equals
	Halt OpCode = 99
)

type OpFunc func(ExecutionContext, ...Variable) OpResult

type OpResult struct {
	halt       bool
	jumpToAddr int
	jump       bool
}

type Operation struct {
	opCode OpCode
	run    OpFunc
	size   int
}

var OpMap = map[OpCode]Operation{
	Add: {
		opCode: Add,
		run:    add,
		size:   3,
	},
	Multiply: {
		opCode: Multiply,
		run:    multiply,
		size:   3,
	},
	ReadStdIn: {
		opCode: ReadStdIn,
		run:    readStdIn,
		size:   1,
	},
	WriteStdOut: {
		opCode: WriteStdOut,
		run:    writeStdOut,
		size:   1,
	},
	JmpIfTrue: {
		opCode: JmpIfTrue,
		run:    jmpIfTrue,
		size:   2,
	},
	JmpIfFalse: {
		opCode: JmpIfFalse,
		run:    jmpIfFalse,
		size:   2,
	},
	LessThan: {
		opCode: LessThan,
		run:    lessThan,
		size:   3,
	},
	Equals: {
		opCode: Equals,
		run:    equals,
		size:   3,
	},
	Halt: {
		opCode: Equals,
		run:    halt,
		size:   0,
	},
}

func halt(ExecutionContext, ...Variable) OpResult {
	return OpResult{halt: true}
}

func add(ctx ExecutionContext, params ...Variable) OpResult {
	param1 := params[0]
	param2 := params[1]
	output := params[2]

	output.Set(param1.Get() + param2.Get())
	return OpResult{}
}

func multiply(ctx ExecutionContext, params ...Variable) OpResult {
	param1 := params[0]
	param2 := params[1]
	output := params[2]

	output.Set(param1.Get() * param2.Get())
	return OpResult{}
}

func readStdIn(ctx ExecutionContext, params ...Variable) OpResult {
	output := params[0]

	userInput := ctx.stdIn.Read()
	output.Set(userInput)

	return OpResult{}
}

func writeStdOut(ctx ExecutionContext, params ...Variable) OpResult {
	param1 := params[0]

	outputValue := param1.Get()
	ctx.stdOut.Write(outputValue)

	return OpResult{}
}

func jmpIfTrue(ctx ExecutionContext, params ...Variable) OpResult {
	param1 := params[0]
	jumpAddress := params[1]

	opResult := OpResult{}
	if param1.Get() != 0 {
		opResult.jump = true
		opResult.jumpToAddr = jumpAddress.Get()
	}
	return opResult
}

func jmpIfFalse(ctx ExecutionContext, params ...Variable) OpResult {
	param1 := params[0]
	jumpAddress := params[1]

	opResult := OpResult{}
	if param1.Get() == 0 {
		opResult.jump = true
		opResult.jumpToAddr = jumpAddress.Get()
	}
	return opResult
}

func lessThan(ctx ExecutionContext, params ...Variable) OpResult {
	param1 := params[0]
	param2 := params[1]
	output := params[2]

	isLessThan := param1.Get() < param2.Get()

	output.Set(utils.BoolToInt(isLessThan))
	return OpResult{}
}

func equals(ctx ExecutionContext, params ...Variable) OpResult {
	param1 := params[0]
	param2 := params[1]
	output := params[2]

	isEqual := param1.Get() == param2.Get()

	output.Set(utils.BoolToInt(isEqual))
	return OpResult{}
}
