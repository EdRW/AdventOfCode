package machine

import (
	"aoc/utils"
	"fmt"
)

type OpCode int

const (
	Add OpCode = iota + 1
	Multiply
	ReadStdIn
	WriteStdOut
	Halt OpCode = 99
)

type OpFunc func(...Variable)

type Operation struct {
	opCode    OpCode
	run       OpFunc
	numParams int
}

var OpMap = map[OpCode]Operation{
	Add: {
		opCode:    Add,
		run:       add,
		numParams: 3,
	},
	Multiply: {
		opCode:    Multiply,
		run:       multiply,
		numParams: 3,
	},
	ReadStdIn: {
		opCode:    ReadStdIn,
		run:       readStdIn,
		numParams: 1,
	},
	WriteStdOut: {
		opCode:    WriteStdOut,
		run:       writeStdOut,
		numParams: 1,
	},
	Halt: noOp(Halt),
}

func noOp(opCode OpCode) Operation {
	return Operation{
		opCode,
		func(params ...Variable) {},
		0,
	}
}

func add(params ...Variable) {
	param1 := params[0]
	param2 := params[1]
	output := params[2]

	output.Set(param1.Get() + param2.Get())
}

func multiply(params ...Variable) {
	param1 := params[0]
	param2 := params[1]
	output := params[2]

	output.Set(param1.Get() * param2.Get())
}

func readStdIn(params ...Variable) {
	output := params[0]

	fmt.Print("Enter the System ID: ")
	var userInput int
	utils.OrDie(fmt.Scanf("%d", &userInput))

	output.Set(userInput)
}

func writeStdOut(params ...Variable) {
	param1 := params[0]

	fmt.Printf("Output: %d\n", param1.Get())
}
