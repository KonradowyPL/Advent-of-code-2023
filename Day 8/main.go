package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Node struct {
	left  string
	right string
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

func findLCM(nums []int) int {
	if len(nums) == 0 {
		return 0 // Return 0 for an empty slice or handle differently if required
	}

	result := nums[0]
	for i := 1; i < len(nums); i++ {
		result = lcm(result, nums[i])
	}
	return result
}

func partBsolve(pos string, nodes *map[string]Node, dirList string) int {
	time := 0

	for i := 0; string(pos[2]) != "Z"; i++ {
		dir := string(dirList[i%len(dirList)])
		node := (*nodes)[pos]

		if dir == "L" {
			pos = node.left
		} else {
			pos = node.right
		}

		time++
	}

	return time
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var lines []string = strings.Split(string(content), "\r\n")

	var poses []string

	nodes := make(map[string]Node)

	var dirList string = lines[0]

	for _, line := range lines[2:] {

		data := strings.Fields(line)
		name := data[0]

		temp := data[2][1:]
		left := temp[:len(temp)-1]

		temp = data[3]
		right := temp[:len(temp)-1]

		node := Node{left, right}

		nodes[name] = node

		if string(name[2]) == "A" {
			poses = append(poses, name)
		}
	}

	time := 0
	pos := "AAA"

	for i := 0; pos != "ZZZ"; i++ {
		dir := string(dirList[i%len(dirList)])
		node := nodes[pos]

		if dir == "L" {
			pos = node.left
		} else {
			pos = node.right

		}

		time++
	}

	fmt.Println(time)

	var times []int

	for _, pos := range poses {
		time := partBsolve(pos, &nodes, dirList)
		times = append(times, time)
	}

	fullTime := findLCM(times)
	fmt.Println(fullTime)
}
