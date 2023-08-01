package main

import (
	"fmt"
	"log"
	"os"
)

type pair struct {
	x int
	y int
}

func move(c byte, x int, y int) (int, int) {
	switch c {
	case '>':
		x += 1
	case '^':
		y += 1
	case 'v':
		y += -1
	case '<':
		x += -1
	}
	return x, y
}

func part1(f []byte) int {
	grid := make(map[pair]int)
	x := 0
	y := 0
	count := 0
	for i := range f {
		grid[pair{x, y}] += 1
		if grid[pair{x, y}] == 1 {
			count += 1
		}
		x, y = move(f[i], x, y)
	}
	return count
}

func part2(f []byte) int {
	grid := make(map[pair]int)
	x1 := 0
	y1 := 0
	x2 := 0
	y2 := 0
	count := 1
	grid[pair{0, 0}] = 1
	for i := 0; i < len(f); i += 2 {
		x1, y1 = move(f[i], x1, y1)
		x2, y2 = move(f[i+1], x2, y2)
		grid[pair{x1, y1}] += 1
		if grid[pair{x1, y1}] == 1 {
			count += 1
		}
		grid[pair{x2, y2}] += 1
		if grid[pair{x2, y2}] == 1 {
			count += 1
		}
	}
	return count
}

func main() {
	f, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Part 1 =", part1(f))
	fmt.Println("Part 1 =", part2(f))
}
