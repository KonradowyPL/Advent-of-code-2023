package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func solve(line []int) (int, int) {
	previous := line[0]
	all := true
	nums := make([]int, len(line)-1)

	for index, num := range line[1:] {
		change := num - previous

		nums[index] = change

		if change != 0 {
			all = false
		}
		previous = num
	}

	if !all {
		last, first := solve(nums)
		return line[len(line)-1] + last, line[0] - first
	} else {
		return line[0], line[len(line)-1]
	}

}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var lines []string = strings.Split(string(content), "\r\n")

	var sum int = 0
	var sumPrevious int = 0

	for _, line := range lines {
		numsTemp := strings.Fields(line)
		nums := make([]int, len(numsTemp))

		for index, _num := range numsTemp {
			num, _ := strconv.Atoi(_num)
			nums[index] = num
		}
		last, first := solve(nums)
		sum += last
		sumPrevious += first
	}
	fmt.Println("next vaues sum:", sum)
	fmt.Println("previous vaues sum:", sumPrevious)

}
