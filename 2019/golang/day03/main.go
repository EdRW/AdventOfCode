package main

import (
	"aoc/utils"
	"fmt"
	"strings"
)

type Point struct {
	X int
	Y int
}

func NewPoint(x int, y int) Point {
	return Point{x, y}
}

type WirePathSteps map[Point]int

func getWirePathSteps(pathString []string) WirePathSteps {
	wirePath := make(WirePathSteps)

	xyCoords := NewPoint(0, 0)
	step := 0
	for _, event := range pathString {
		direction := event[0]
		magnitude := utils.ToInt(event[1:])

		sign := 1
		if direction == 'L' || direction == 'D' {
			sign = -1
		}

		for i := 0; i < magnitude; i++ {
			if direction == 'L' || direction == 'R' {
				xyCoords.X += sign
			} else {

				xyCoords.Y += sign
			}
			step += 1
			wirePath[NewPoint(xyCoords.X, xyCoords.Y)] = step
		}
	}

	return wirePath
}

func Abs(number int) int {
	if number < 0 {
		return number * -1
	}
	return number
}

func mhtnDistance(point Point) int {
	return Abs(point.X) + Abs(point.Y)
}

func stepDistance(point Point,
	wirePathStepsA WirePathSteps,
	wirePathStepsB WirePathSteps) int {
	stepsA := wirePathStepsA[point]
	stepsB := wirePathStepsB[point]
	return stepsA + stepsB
}

func main() {
	input := utils.AocInputFile(3)
	scanner, close := utils.NewFileScanner(input)
	defer close()

	wirePaths := make([]WirePathSteps, 2)
	for i := 0; scanner.Scan(); i++ {
		inputTxt := scanner.Text()
		pathString := strings.Split(inputTxt, ",")
		wirePaths[i] = getWirePathSteps(pathString)
	}

	wirePath1 := wirePaths[0]
	wirePath2 := wirePaths[1]

	intersections := make([]Point, 0)

	for k := range wirePath1 {
		if _, ok := wirePath2[k]; ok {
			intersections = append(intersections, k)
		}
	}

	if len(intersections) == 0 {
		fmt.Println("The paths don't cross!")
		return
	}

	firstIntersection := intersections[0]

	leastDistance := mhtnDistance(firstIntersection)
	leastSteps := stepDistance(firstIntersection, wirePath1, wirePath2)

	for _, intersection := range intersections[1:] {
		distance := mhtnDistance(intersection)
		if distance < leastDistance {
			leastDistance = distance
		}

		stepDistance := stepDistance(intersection, wirePath1, wirePath2)
		if stepDistance < leastSteps {
			leastSteps = stepDistance
		}
	}

	fmt.Printf("Part 1: closest intersection is %d blocks away!\n", leastDistance)
	fmt.Printf("Part 2: least steps to intersection is %d steps away!\n", leastSteps)
}
