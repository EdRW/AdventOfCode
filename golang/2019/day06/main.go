package main

import (
	"aoc/utils"
	"fmt"
	"strings"
)

type TotalOrbitsMap = map[string]int
type OrbitalMap = map[string]utils.Set[string]

func countSatelliteOrbits(object string, orbitalMap OrbitalMap, numOrbitsMap TotalOrbitsMap) {
	satellites := orbitalMap[object]
	numOrbits := numOrbitsMap[object]

	for satellite := range satellites {
		numOrbitsMap[satellite] = numOrbits + 1
		countSatelliteOrbits(satellite, orbitalMap, numOrbitsMap)
	}
}

func searchSatelliteOrbits(target string, root string, orbitalMap OrbitalMap) []string {

	if orbitalMap[root].Exists(target) {
		return []string{root}
	}

	for satellite := range orbitalMap[root] {
		path := searchSatelliteOrbits(target, satellite, orbitalMap)
		if len(path) > 0 {
			return append([]string{root}, path...)
		}
	}

	return []string{}
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
			existingSatellites.Add(satellite)
			orbitalMap[object] = existingSatellites
		} else {
			satellites := utils.NewSet(satellite)
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

	fmt.Println()

	fmt.Println("~~~~~ Part 2 ~~~~~")
	// make arrays describing the paths between COM and YOU
	// and between COM and SANTA
	// traverse both arrays until they diverge.
	// then add the remaining len of the arrays
	youPath := searchSatelliteOrbits("YOU", COM, orbitalMap)
	santaPath := searchSatelliteOrbits("SAN", COM, orbitalMap)
	fmt.Println(youPath)
	fmt.Println(santaPath)
	i := 0
	for i < utils.Min(len(youPath), len(santaPath)) {
		if youPath[i] != santaPath[i] {
			break
		}
		i++
	}
	fmt.Printf("i: %d\n", i)
	minJumps := len(youPath) - i + len(santaPath) - i

	fmt.Printf("Min Jumps: %d", minJumps)
}
