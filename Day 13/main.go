package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func rotate(matrix [][]byte) [][]byte {

	newMatrix := make([][]byte, len(matrix[0]))
	for i := 0; i < len(newMatrix); i++ {
		newMatrix[i] = make([]byte, len(matrix))
	}

	for i := 0; i < len(newMatrix); i++ {
		for j := 0; j < len(newMatrix[0]); j++ {
			newMatrix[i][j] = matrix[j][i]

		}

	}

	return newMatrix
}

func comfirm(lines []string, index int) bool {
	checks := min(index, len(lines)-index)
	for i := 0; i < checks; i++ {

		if lines[index-i-1] != lines[index+i] {
			return false
		}
	}

	return true
}

func getHorisontal(lines []string) int {
	previous := ""
	for i, line := range lines {
		if line == previous {
			if comfirm(lines, i) {
				return i
			}
		}
		previous = line
	}

	return 0
}

func getVertical(lines []string) int {
	var byteArr [][]byte
	for _, line := range lines {
		byteArr = append(byteArr, []byte(line))
	}

	byteArr = rotate(byteArr)
	var _lines []string

	for _, line := range byteArr {
		_lines = append(_lines, string(line))
	}

	return getHorisontal(_lines)

}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var mirrors []string = strings.Split(string(content), "\r\n\r\n")

	sum := 0

	for _, mirror := range mirrors {
		horisontal := getHorisontal(strings.Split(mirror, "\r\n"))
		vertical := getVertical(strings.Split(mirror, "\r\n"))

		sum += horisontal*100 + vertical
	}
	fmt.Println(sum)
}
