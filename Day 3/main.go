package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func checkPart(row int, col int, lines *[]string) bool {

	if row < 0 || row >= len(*lines) {
		return false
	}
	if col < 0 || col >= len((*lines)[row]) {
		return false
	}

	var char = (*lines)[row][col]
	return char != 46 && !(57 >= char && char >= 48)
}

func checkNum(row int, col int, lines *[]string) bool {
	if row < 0 || row >= len(*lines) {
		return false
	}
	if col < 0 || col >= len((*lines)[row]) {
		return false
	}

	var char = (*lines)[row][col]
	return (57 >= char && char >= 48)
}

func findNumAt(row int, col int, lines *[]string) int {
	if row < 0 || row >= len(*lines) {
		return 0
	}
	if col < 0 || col >= len((*lines)[row]) {
		return 0
	}
	for checkNum(row, col, lines) {
		col--
	}
	col++
	var numbers []byte
	for checkNum(row, col, lines) {
		numbers = append(numbers, (*lines)[row][col])
		col++
	}

	num, _ := strconv.Atoi(string(numbers))

	return num
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var lines []string = strings.Split(string(content), "\r\n")

	var sum int = 0
	var ratioSum int = 0

	for lineNum, line := range lines {
		line = line + "."
		var num []rune = nil
		var isPart bool = false

		for charIdx, char := range line {

			if 57 >= char && char >= 48 {
				num = append(num, char)

				var checks [][2]int = [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}

				//check if is part
				for _, check := range checks {
					x := check[0]
					y := check[1]
					isPart = isPart || checkPart(lineNum+x, charIdx+y, &lines)
				}

			} else {
				if isPart {
					number, _ := strconv.Atoi(string(num))
					sum += number
				}
				num = nil
				isPart = false
			}
			if char == 42 {
				var checks [][2]int = [][2]int{
					{-1, -1}, {-1, 0}, {-1, 1},

					{0, -1},

					{0, 1},

					{1, -1}, {1, 0}, {1, 1},
				}

				var isGear int = 0
				var values []int
				for _, check := range checks {
					x := check[0]
					y := check[1]

					var thisNum bool = checkNum(lineNum+x, charIdx+y, &lines)

					if thisNum {
						num := findNumAt(lineNum+x, charIdx+y, &lines)

						found := false
						for _, v := range values {
							if v == num {
								found = true
								break
							}
						}
						if !found {
							values = append(values, num)
							isGear++

						}

					}

				}

				if isGear == 2 {

					var ratio int = values[0] * values[1]
					ratioSum += ratio
				}
			}

		}
	}
	fmt.Println("sum", sum)
	fmt.Println("ratio sum", ratioSum)
}
