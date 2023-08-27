package main

import (
	"aoc/utils"
	"fmt"
	"strings"
)

type TotalOrbitsMap = map[string]int
type OrbitalMap = map[string][]string

func countSatelliteOrbits(object string, orbitalMap OrbitalMap, numOrbitsMap TotalOrbitsMap) {
	satellites := orbitalMap[object]
	numOrbits := numOrbitsMap[object]

	for _, satellite := range satellites {
		numOrbitsMap[satellite] = numOrbits + 1
		countSatelliteOrbits(satellite, orbitalMap, numOrbitsMap)
	}
}

func readOrbitalMap() OrbitalMap {

	input := utils.AOCInputFile(6)

	scanner, close := utils.NewFileScanner(input)
	defer close()

	orbitalMap := OrbitalMap{}
	for scanner.Scan() {
		inputTxt := scanner.Text()
		objectAndSatellite := strings.Split(inputTxt, ")")

		object := objectAndSatellite[0]
		satellite := objectAndSatellite[1]

		if existingSatellites, ok := orbitalMap[object]; ok {
			orbitalMap[object] = append(existingSatellites, satellite)
		} else {
			satellites := []string{satellite}
			orbitalMap[object] = satellites
		}
	}
	return orbitalMap
}

func main() {
	orbitalMap := readOrbitalMap()
	// fmt.Printf("%v", orbitalMap)

	COM := "COM"
	numOrbitsMap := TotalOrbitsMap{COM: 0}

	// iterate over the objects in orbitalMap starting at COM
	// count the orbits and store them in numOrbitsMap
	countSatelliteOrbits(COM, orbitalMap, numOrbitsMap)

	// count the total number of orbits
	totalOrbits := 0
	for _, numOrbits := range numOrbitsMap {
		totalOrbits += numOrbits
	}

	fmt.Printf("Total Orbits: %d", totalOrbits)
}
