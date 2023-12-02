package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var lines []string = strings.Split(string(content), "\r\n")

	const maxRed int = 12
	const maxGreen int = 13
	const maxBlue int = 14

	var idsSum int = 0
	var powerSum int = 0

	for _, line := range lines {
		var isPossible bool = true

		var dat []string = strings.Split(line, ": ")
		id, _ := strconv.Atoi(dat[0][5:])

		var cubeSet []string = strings.Split(dat[1], "; ")

		var minRed int = 0
		var minGreen int = 0
		var minBlue int = 0

		for _, cubes := range cubeSet {
			var red = 0
			var green = 0
			var blue = 0

			var cubeList []string = strings.Split(cubes, ", ")
			for _, cube := range cubeList {
				var cubeData []string = strings.Split(cube, " ")

				amount, _ := strconv.Atoi(cubeData[0])
				var color string = cubeData[1]

				switch color {
				case "red":
					red += amount
				case "green":
					green += amount
				case "blue":
					blue += amount
				}

			}

			if red > minRed {
				minRed = red
			}
			if green > minGreen {
				minGreen = green
			}
			if blue > minBlue {
				minBlue = blue
			}

			if red > maxRed || green > maxGreen || blue > maxBlue {
				isPossible = false
			}
		}

		var power int = minRed * minGreen * minBlue

		powerSum += power

		if isPossible {
			idsSum += id
		}
	}
	fmt.Println("IDs sum:", idsSum)
	fmt.Println("powers sum:", powerSum)

}
