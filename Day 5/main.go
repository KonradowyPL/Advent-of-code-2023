package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type seedMap struct {
	name   string
	ranges []seedRange
}

type seedRange struct {
	destonation int
	source      int
	lenght      int
}

func CalculateNextLocation(seed int, _map seedMap) int {
	destonation := seed

	for _, _seedRange := range _map.ranges {
		from := _seedRange.source
		to := from + _seedRange.lenght

		if seed >= from && seed <= to {
			change := seed - from
			destonation = _seedRange.destonation + change

		}
	}

	return destonation
}

func evalSeed(seed int, maps []seedMap) int {
	destonation := seed
	for _, _map := range maps {
		destonation = CalculateNextLocation(destonation, _map)
	}
	return destonation

}

func main() {

	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var lines []string = strings.Split(string(content), "\r\n")

	var seedsTemp []string = strings.Split(strings.Split(lines[0], ": ")[1], " ")
	var seeds []int

	for _, num := range seedsTemp {
		seed, _ := strconv.Atoi(num)
		seeds = append(seeds, seed)

	}

	var maps []seedMap

	for _, _map := range strings.Split(string(content), "\r\n\r\n") {
		mapLines := strings.Split(_map, "\r\n")

		name := mapLines[0]

		var ranges []seedRange

		for _, seedRanges := range mapLines[1:] {
			rangesData := strings.Split(seedRanges, " ")

			destonation, _ := strconv.Atoi(rangesData[0])
			source, _ := strconv.Atoi(rangesData[1])
			lenght, _ := strconv.Atoi(rangesData[2])

			ranges = append(ranges, seedRange{destonation, source, lenght})
		}
		maps = append(maps, seedMap{name, ranges})
	}

	minValA := math.MaxInt
	minValB := math.MaxInt

	for _, seed := range seeds {
		destonation := evalSeed(seed, maps)

		if destonation < minValA {
			minValA = destonation
		}
	}

	fmt.Println("min distance part A", minValA)
	fmt.Println("Warning: Unoptimized code: this will around 5 min")

	for i := 0; i < len(seeds); i += 2 {
		from, _ := strconv.Atoi(seedsTemp[i])
		amount, _ := strconv.Atoi(seedsTemp[i+1])

		fmt.Println("computing seed range:", i>>1)

		for i := from; i < from+amount; i++ {

			destonation := evalSeed(i, maps)

			if destonation < minValB {
				minValB = destonation
			}
		}
	}
	fmt.Println("min distance part B", minValB)

}
