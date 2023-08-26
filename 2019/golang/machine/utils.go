package machine

import (
	"aoc/utils"
	"log"
	"strings"
)

func GetIntCodesFromFile(filePath string) []int {
	scanner, close := utils.NewFileScanner(filePath)
	if !scanner.Scan() {
		log.Fatal()
	}
	close()

	inputTxt := scanner.Text()
	intCodeStrs := strings.Split(inputTxt, ",")
	return utils.ToInts(intCodeStrs)
}

func parseOpCode(opCodeIntCode int) (OpCode, []int) {
	ints := utils.IntToInts(opCodeIntCode)

	opCodeInt := 0
	paramModes := make([]int, 0)

	multiplier := 1
	for i := len(ints) - 1; i >= 0; i-- {
		// opCodes from just the last 2 numbers
		if i >= len(ints)-2 {
			opCodeInt += ints[i] * multiplier
			multiplier *= 10
			continue
		}

		// params from all numbers except the last 2
		paramModes = append(paramModes, ints[i])
	}

	opCode := OpCode(opCodeInt)

	return opCode, paramModes
}
