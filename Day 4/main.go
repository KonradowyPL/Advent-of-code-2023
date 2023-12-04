package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

var deck []int

func removeDuplicates(numbers []string) []string {
	seen := make(map[string]struct{}, len(numbers))
	j := 0
	for _, v := range numbers {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		numbers[j] = v
		j++
	}
	return numbers[:j]
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var lines []string = strings.Split(string(content), "\r\n")

	for i := 0; i < len(lines); i++ {
		deck = append(deck, 1)
	}

	var sumPoints = 0
	var cards = 0

	for index, line := range lines {
		cardData := strings.ReplaceAll(strings.ReplaceAll(strings.Split(line, ":")[1], "| ", ""), "  ", " ")
		numbers := strings.Split(cardData, " ")

		var lenght int = len(numbers)
		numbers = removeDuplicates(numbers)

		var newLen int = len(numbers)
		var change int = lenght - newLen

		var points int = 0
		var amount int = deck[index]

		if change > 0 {
			points = int(math.Pow(2, float64(change)-1))
		}

		for i := 0; i < change; i++ {
			deck[index+i+1] += amount
		}

		sumPoints += points
		cards += amount
	}

	fmt.Println(sumPoints)
	fmt.Println(cards)

}
