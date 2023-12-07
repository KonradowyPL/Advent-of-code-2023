package main

//253926070
//253926070

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type card struct {
	power int
	bid   int
	val   string
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func SliceIndex(value string, list *[]string) int {
	for index, val := range *list {
		if val == value {
			return index
		}
	}
	return -1
}

func getCardPoints(hands string, isPartB bool) (int, int) {
	var data = make(map[rune]int, 5)

	var powers []int = []int{371293, 28561, 2197, 169, 13}
	var digits []string = []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}

	if isPartB {
		digits = []string{"J", "2", "3", "4", "5", "6", "7", "8", "9", "T", "Q", "K", "A"}
	}
	strenght := 0

	for index, hand := range hands {
		if data[hand] == 0 {
			data[hand] = 0
		}
		data[hand]++

		power := (SliceIndex(string(hand), &digits) + 1) * powers[index]

		strenght += power
	}

	jockers := 0

	if isPartB {
		//"J" == 74
		jockers = data[74]

		delete(data, 74)
	}
	repeat := make([]int, 0, len(data))
	for _, value := range data {
		repeat = append(repeat, value)
	}

	sort.Ints(repeat)
	if isPartB {
		if len(repeat) != 0 {
			repeat[len(repeat)-1] += jockers
		} else {
			return 7, strenght
		}
	}

	var p7 []int = []int{5}
	var p6 []int = []int{1, 4}
	var p5 []int = []int{2, 3}
	var p4 []int = []int{1, 1, 3}
	var p3 []int = []int{1, 2, 2}
	var p2 []int = []int{1, 1, 1, 2}

	if Equal(repeat, p7) {
		return 7, strenght
	} else if Equal(repeat, p6) {
		return 6, strenght
	} else if Equal(repeat, p5) {
		return 5, strenght
	} else if Equal(repeat, p4) {
		return 4, strenght
	} else if Equal(repeat, p3) {
		return 3, strenght
	} else if Equal(repeat, p2) {
		return 2, strenght
	} else {
		return 1, strenght
	}
}

func calculateCards(lines []string, isPartB bool) int {
	var cards []card

	for _, line := range lines {
		cardDat := strings.Fields(line)

		hands := cardDat[0]
		bid, _ := strconv.Atoi(cardDat[1])

		points, power := getCardPoints(hands, isPartB)

		// 13 is number of all diits
		// 7 is number of possible configurations
		// 6 is amount of cards in hand
		// (13^6-1) * 7
		strenght := points*33787656 + power

		cards = append(cards, card{strenght, bid, hands})
	}

	sort.Slice(cards, func(i, j int) bool {
		return cards[i].power < cards[j].power
	})

	bidSum := 0

	for index, card := range cards {
		bidSum += card.bid * (index + 1)
	}

	return bidSum
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var lines []string = strings.Split(string(content), "\r\n")

	fmt.Println("part A sum", calculateCards(lines, false))
	fmt.Println("part B sum", calculateCards(lines, true))
}
