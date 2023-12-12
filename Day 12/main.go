package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func Equal(a []int, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		bi, _ := strconv.Atoi(b[i])
		if v != bi {
			return false
		}
	}
	return true
}

func parse(line string) []int {
	var lens []int
	len := 0
	for _, char := range line {
		if string(char) == "." && len != 0 {
			lens = append(lens, len)
			len = 0
		} else if string(char) == "#" {
			len++
		}
	}
	if len != 0 {
		lens = append(lens, len)
	}

	return lens
}

func getNthBit(val, n uint32) int {
	mask := 1 << (n - 1)
	return int((val & uint32(mask)) >> (n - 1))
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var lines []string = strings.Split(string(content), "\r\n")

	fmt.Println("part A", solve(lines, false))
	fmt.Println("part B", solve(lines, true))

}

func solve(lines []string, partB bool) int {
	posibilities := 0
	for _, line := range lines {

		dat := strings.Fields(line)

		mapDat := dat[0]
		var springsNext []string

		if partB {

			mapDat = strings.Repeat(mapDat+"?", 4) + mapDat

			springsNext = strings.Split(strings.Repeat(dat[1]+",", 4)+dat[1], ",")
		} else {
			springsNext = strings.Split(dat[1], ",")

		}
		amount := strings.Count(mapDat, "?")

		for i := 0; i < int(math.Pow(2, float64(amount))); i++ {
			mapBytes := []byte(mapDat)
			springs := 0

			for index, char := range mapBytes {
				if string(char) == "?" {

					if getNthBit(uint32(i), uint32(springs+1)) == 0 {
						mapBytes[index] = []byte(".")[0]
					} else {
						mapBytes[index] = []byte("#")[0]
					}

					springs++
				}
			}

			lakes := parse(string(mapBytes))

			if Equal(lakes, springsNext) {
				posibilities++
			}
		}
	}
	return posibilities
}
