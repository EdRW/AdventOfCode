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

func OrDie1[T any](err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func AOCDir(dayNum int) string {
	return fmt.Sprintf("day%02d", dayNum)
}

func AOCInputFile(dayNum int) string {
	return fmt.Sprintf("%s/input.txt", AOCDir(dayNum))
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

func ToString(number int) string {
	return strconv.Itoa(number)
}

func ToStrings(ints []int) []string {
	strs := make([]string, len(ints))
	for i, number := range ints {
		strs[i] = ToString(number)
	}
	return strs
}

func NumDigits(num int) int {
	numDigits := 1

	multiplier := 10
	for num%multiplier < num {
		numDigits++
		multiplier *= 10
	}

	return numDigits
}

func IntToInts(num int) []int {
	digits := make([]int, NumDigits(num))

	for i := len(digits) - 1; num > 0; i-- {
		digits[i] = num % 10
		num = num / 10
	}
	return digits
}

func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func Min(num1 int, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}

func Max(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}

func Factorial(n int) int {
	result := 1
	for i := n; i > 0; i-- {
		result *= i
	}
	return result
}

func Copy[T any](slice []T) []T {
	sliceCopy := make([]T, len(slice))
	copy(sliceCopy, slice)
	return sliceCopy
}

// Prepends elems... to the front of slice
func Prepend[T any](slice []T, elems ...T) []T {
	return append(elems, slice...)
}

func All[T any](slice []T, pred func(T) bool) bool {
	for _, elem := range slice {
		if !pred(elem) {
			return false
		}
	}
	return true
}
