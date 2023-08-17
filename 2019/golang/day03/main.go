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

type WirePath struct {
	set utils.Set[Point]
}

func (w WirePath) Add(point Point) {
	w.set.Add(point)
}

func (w WirePath) Exists(point Point) bool {
	return w.set.Exists(point)
}

func (w WirePath) Range() utils.Set[Point] {
	return w.set
}

func NewWirePath() WirePath {
	return WirePath{utils.NewSet[Point]()}
}

func getWirePath(pathString []string) WirePath {
	var wirePath WirePath = NewWirePath()

	xyCoords := NewPoint(0, 0)

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
			wirePath.Add(NewPoint(xyCoords.X, xyCoords.Y))
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

func main() {
	input := utils.AocInputFile(3)
	scanner, close := utils.NewFileScanner(input)
	defer close()

	wirePaths := make([]WirePath, 2)
	for i := 0; scanner.Scan(); i++ {
		inputTxt := scanner.Text()
		pathString := strings.Split(inputTxt, ",")
		wirePaths[i] = getWirePath(pathString)
	}

	wirePath1 := wirePaths[0]
	wirePath2 := wirePaths[1]

	intersections := make([]Point, 0)

	for k := range wirePath1.Range() {
		if ok := wirePath2.Exists(k); ok {
			intersections = append(intersections, k)
		}
	}

	if len(intersections) == 0 {
		fmt.Println("The paths don't cross!")
		return
	}

	leastDistance := mhtnDistance(intersections[0])
	for _, intersection := range intersections[1:] {
		distance := mhtnDistance(intersection)
		if distance < leastDistance {
			leastDistance = distance
		}
	}

	fmt.Printf("closest intersection is %d blocks away!\n", leastDistance)
}
