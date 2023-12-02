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
	var suma int = 0
	var sumb int = 0

	dict := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}

	for _, line := range lines {

		var digitsa []rune
		var digitsb []rune

		var digit []rune

		for _, char := range line {
			if 57 >= char && char >= 48 {
				digitsa = append(digitsa, char)
				digitsb = append(digitsb, char)

			} else {
				digit = append(digit, char)
				for index, item := range dict {

					if strings.HasSuffix(string(digit), item) {
						digitsb = append(digitsb, rune(strconv.Itoa(index + 1)[0]))
					}
				}
			}
		}
		var spliteda string = string(digitsa[0]) + string(digitsa[len(digitsa)-1])
		var splitedb string = string(digitsb[0]) + string(digitsb[len(digitsb)-1])
		valuea, _ := strconv.Atoi(spliteda)
		valueb, _ := strconv.Atoi(splitedb)

		suma += valuea
		sumb += valueb
	}

	fmt.Println("Sum a:", suma)
	fmt.Println("Sum b:", sumb)

}
