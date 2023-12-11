package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Galaxy struct {
	x int
	y int
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func solve(multipler int, lines []string, galaxies []Galaxy) int {
	fakeY := 0
	fakeX := 0

	for _, line := range lines {
		fakeX = 0

		if string(line[0]) == "-" {
			fakeY += multipler

		} else if string(line[0]) == "@" {
			fakeY += multipler
			fakeX += multipler

		}

		for _, char := range line {

			if string(char) == "#" {
				galaxies = append(galaxies, Galaxy{fakeX, fakeY})
			} else if string(char) == "|" {
				fakeX += multipler
			}

			fakeX++
		}
		fakeY++

	}

	sum := 0

	for i, galaxyA := range galaxies {
		for _, galaxyB := range galaxies[i+1:] {
			x := galaxyA.x - galaxyB.x
			y := galaxyA.y - galaxyB.y

			sum += abs(x) + abs(y)
		}
	}

	return sum
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var lines []string = strings.Split(string(content), "\r\n")

	fullCols := make([]bool, len(lines[0]))

	for index := 0; index < len(lines); index++ {
		line := lines[index]

		empty := true
		for charIdx, char := range line {
			if string(char) == "#" {
				empty = false
				fullCols[charIdx] = true
			}
		}
		if empty {
			lines = append(lines[:index+1], lines[index:]...)
			lines[index] = strings.Repeat("-", len(lines[0]))
			index++
		}
	}

	for index, line := range lines {
		realIdx := 0
		_bytes := []byte(line)
		for charIdx := 0; realIdx < len(fullCols); charIdx++ {
			if !fullCols[realIdx] {

				_bytes = append(_bytes[:charIdx+1], _bytes[charIdx:]...)

				if _bytes[charIdx] == []byte("-")[0] {
					_bytes[charIdx] = []byte("@")[0]
				} else {
					_bytes[charIdx] = []byte("|")[0]
				}

				charIdx++
			}
			realIdx++
		}
		lines[index] = string(_bytes)
	}

	var galaxies []Galaxy

	fmt.Println("part a:", solve(0, lines, galaxies))
	fmt.Println("part b:", solve(1000000-2, lines, galaxies))

}
