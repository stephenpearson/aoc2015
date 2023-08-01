package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type pair struct {
	x int
	y int
}

func splitPair(field string) (int, int) {
	params := strings.Split(field, ",")
	v1, _ := strconv.Atoi(params[0])
	v2, _ := strconv.Atoi(params[1])
	return v1, v2
}

func apply(grid *map[pair]int, x1 int, y1 int, x2 int, y2 int, f func(val int) int) {
	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			p := pair{i, j}
			(*grid)[p] = f((*grid)[p])
		}
	}
}

func run(filename string, fn map[string]func(int) int) int {
	grid := make(map[pair]int)

	f, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(f), "\n")

	for i := range lines {
		fields := strings.Split(lines[i], " ")
		if len(fields) < 4 {
			continue
		}
		if fields[0] == "toggle" {
			x1, y1 := splitPair(fields[1])
			x2, y2 := splitPair(fields[3])
			apply(&grid, x1, y1, x2, y2, fn["toggle"])
		} else {
			x1, y1 := splitPair(fields[2])
			x2, y2 := splitPair(fields[4])
			if fields[1] == "on" {
				apply(&grid, x1, y1, x2, y2, fn["turn_on"])
			} else {
				apply(&grid, x1, y1, x2, y2, fn["turn_off"])
			}
		}
	}
	count := 0
	for i := range grid {
		count += grid[i]
	}
	return count
}

func main() {
	part1_fns := map[string]func(int) int{
		"toggle": func(v int) int {
			if v == 0 {
				return 1
			} else {
				return 0
			}
		},
		"turn_on": func(v int) int {
			return 1
		},
		"turn_off": func(v int) int {
			return 0
		},
	}

	part2_fns := map[string]func(int) int{
		"toggle": func(v int) int {
			return v + 2
		},
		"turn_on": func(v int) int {
			return v + 1
		},
		"turn_off": func(v int) int {
			if v > 0 {
				return v - 1
			} else {
				return 0
			}
		},
	}

	fmt.Println("Part 1 =", run("input", part1_fns))
	fmt.Println("Part 2 =", run("input", part2_fns))
}
