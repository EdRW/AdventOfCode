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

var OpFuncMap = map[OpCode]Operation{
	Add:         addOp,
	Multiply:    multiplyOp,
	ReadStdIn:   readStdInOp,
	WriteStdOut: writeStdOutOp,
	Halt:        noOp(Halt),
}

var addOp = Operation{
	Add, add, 3,
}

var multiplyOp = Operation{
	Multiply, multiply, 3,
}

var readStdInOp = Operation{
	ReadStdIn, readStdIn, 1,
}

var writeStdOutOp = Operation{
	WriteStdOut, writeStdOut, 1,
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
