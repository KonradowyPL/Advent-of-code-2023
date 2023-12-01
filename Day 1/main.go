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
	var sum int = 0

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

		var digits []rune
		var digit []rune

		for _, char := range line {
			if 57 >= char && char >= 48 {
				digits = append(digits, char)

			} else {
				digit = append(digit, char)
				for index, item := range dict {

					if strings.HasSuffix(string(digit), item) {
						digits = append(digits, rune(strconv.Itoa(index + 1)[0]))
					}
				}
			}
		}

		var splited string = string(digits[0]) + string(digits[len(digits)-1])
		value, _ := strconv.Atoi(splited)

		sum += value
	}

	fmt.Println(sum)
}
