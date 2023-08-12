package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func OrDie[T any](val T, err error) T {
	if err != nil {
		log.Fatal(err)
	}
	return val
}

func AocDir(dayNum int) string {
	return fmt.Sprintf("day%02d", dayNum)
}
func AocInputFile(dayNum int) string {
	return fmt.Sprintf("%s/input.txt", AocDir(dayNum))
}

// NewFileScanner opens a file and
// returns an associated *bufio.Scanner
// as well as a func to close the file.
//
//	scanner, close := NewFileScanner("day01/input.txt")
//	defer close()
func NewFileScanner(inputPath string) (scanner *bufio.Scanner, close func() error) {
	file := OrDie(os.Open(inputPath))
	return bufio.NewScanner(file), file.Close
}

func ToInt(str string) int {
	return OrDie(strconv.Atoi(str))
}

func ToInts(strs []string) []int {
	var ints = make([]int, len(strs))
	for i, str := range strs {
		ints[i] = ToInt(str)

	}
	return ints
}
