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

// ref: idea for switch stolen from Camille.Codes
func (p Point) Move(direction byte, magnitude int) Point {
	switch direction {
	case 'L':
		p.X -= magnitude
	case 'R':
		p.X += magnitude
	case 'U':
		p.Y += magnitude
	case 'D':
		p.Y -= magnitude
	}
	return p
}

type WirePathSteps map[Point]int

func (w WirePathSteps) Intersections(otherWirePath WirePathSteps) []Point {
	intersections := make([]Point, 0)

	for k := range w {
		if _, ok := otherWirePath[k]; ok {
			intersections = append(intersections, k)
		}
	}

	return intersections
}

func getWirePathSteps(pathString []string) WirePathSteps {
	wirePath := make(WirePathSteps)
	point := NewPoint(0, 0)

	step := 0
	for _, event := range pathString {
		direction := event[0]
		magnitude := utils.ToInt(event[1:])

		for i := 0; i < magnitude; i++ {
			point = point.Move(direction, 1)
			step += 1
			wirePath[point] = step
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

func stepDistance(
	point Point,
	wirePathStepsA WirePathSteps,
	wirePathStepsB WirePathSteps,
) int {
	stepsA := wirePathStepsA[point]
	stepsB := wirePathStepsB[point]
	return stepsA + stepsB
}

func main() {
	input := utils.AOCInputFile(3)
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

	intersections := wirePath1.Intersections(wirePath2)

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
