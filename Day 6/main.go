package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func calcSolutions(time int, distance int) int {
	// a little bit of math:
	// t = time for race
	// x = time we are holding down the button
	// distance: d = x * (t - x) = x*t - x^2
	// we have to calculate x from this formula

	// x = (t ± sqrt(t^2 - 4 d))/ 2
	// we can replace ± with + for higher time or
	// we can replace ± with - for lower time

	var preCalc = math.Sqrt(float64(time*time - 4*(distance+1)))

	minTime := math.Ceil((float64(time) - preCalc) * 0.5)
	maxTime := math.Floor((float64(time) + preCalc) * 0.5)

	// subreact higher from lower and add 1 for amount of solutions

	solutions := int(maxTime - minTime + 1)

	return solutions
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var lines []string = strings.Split(string(content), "\r\n")

	var margin int = 1

	times := strings.Fields(lines[0])
	distances := strings.Fields(lines[1])

	for index, _time := range times[1:] {
		time, _ := strconv.Atoi(_time)
		distance, _ := strconv.Atoi(distances[index+1])

		solutions := calcSolutions(time, distance)

		margin *= solutions
	}

	// part b:
	// removing all spaces and making them a single number
	time, _ := strconv.Atoi(strings.ReplaceAll(strings.Split(lines[0], ":")[1], " ", ""))
	distance, _ := strconv.Atoi(strings.ReplaceAll(strings.Split(lines[1], ":")[1], " ", ""))

	solutions := calcSolutions(time, distance)

	fmt.Println("margin:", margin)
	fmt.Println("solutions:", solutions)
}
