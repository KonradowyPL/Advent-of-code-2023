package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type position struct {
	x int
	y int
}

type direction struct {
	x   int
	y   int
	dir int
}

type pipe struct {
	// 1 north
	// 2 east
	// 3 south
	// 4 west
	dir1 int
	dir2 int
}

func getAt(x int, y int, lines *[]string) byte {
	if y >= len(*lines) || y < 0 {
		return 46
	}
	if x >= len((*lines)[y]) || x < 0 {
		return 46
	}

	return (*lines)[y][x]
}

func pointsIn(_pipe pipe, dir int) bool {
	return _pipe.dir1 == dir || _pipe.dir2 == dir
}

func move(startPos position, lines *[]string) {
	pos := startPos

	pipes := map[byte]pipe{
		124: {1, 3},
		45:  {2, 4},
		76:  {1, 2},
		74:  {1, 4},
		70:  {2, 3},
		55:  {3, 4},
		46:  {0, 0},
		83:  {0, 0},
	}

	neighbors := []direction{
		{0, -1, 1}, // north  (up)
		{1, 0, 2},  // east   (right)
		{0, 1, 3},  // south  (down)
		{-1, 0, 4}, // west   (left)
	}

	var poses [2]direction
	posIdx := 0

	for _, neighbor := range neighbors {
		x := pos.x + neighbor.x
		y := pos.y + neighbor.y

		_pipe := getAt(x, y, lines)

		pipeDir := pipes[_pipe]

		points := pointsIn(pipeDir, (neighbor.dir+1)%4+1)

		if points {
			poses[posIdx] = direction{x, y, neighbor.dir}
			posIdx++
		}
	}

	// for i := 0; i < 10; i++ {
	steps := 1
	for ; !(poses[0].x == poses[1].x && poses[0].y == poses[1].y); steps++ {
		for index, pos := range poses {
			_pipe := getAt(pos.x, pos.y, lines)
			pipeDir := pipes[_pipe]
			reverse := ((pos.dir + 1) % 4) + 1
			var moveDir int

			if pipeDir.dir1 != reverse {
				moveDir = pipeDir.dir1
			} else {
				moveDir = pipeDir.dir2
			}
			_move := neighbors[(moveDir+3)%4]
			x := pos.x + _move.x
			y := pos.y + _move.y

			poses[index] = direction{x, y, _move.dir}
		}
	}
	fmt.Println(steps)

}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var lines []string = strings.Split(string(content), "\r\n")

	var pos position = position{0, 0}

	for y, line := range lines {
		for x, char := range line {
			if string(char) == "S" {
				pos = position{x, y}
			}
		}
	}

	move(pos, &lines)

}
