package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func neighbours(grid [][]rune, x int, y int) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			xp := j + x
			yp := i + y
			if yp < 0 || yp >= len(grid) || xp < 0 || xp >= len(grid[yp]) {
				continue
			} else {
				if grid[yp][xp] == '#' {
					count += 1
				}
			}
		}
	}
	return count
}

func cornerLights(grid [][]rune) {
	grid[0][0] = '#'
	grid[len(grid)-1][0] = '#'
	grid[0][len(grid[0])-1] = '#'
	grid[len(grid)-1][len(grid[0])-1] = '#'
}

func iterate(grid [][]rune, corners bool) [][]rune {
	result := [][]rune{}
	for y := 0; y < 100; y++ {
		line := []rune{}
		for x := 0; x < 100; x++ {
			n := neighbours(grid, x, y)
			if grid[y][x] == '#' {
				if n == 2 || n == 3 {
					line = append(line, '#')
				} else {
					line = append(line, '.')
				}
			} else {
				if n == 3 {
					line = append(line, '#')
				} else {
					line = append(line, '.')
				}
			}
		}
		result = append(result, line)
	}
	cornerLights(result)
	return result
}

func read(filename string) [][]rune {
	f, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	grid := [][]rune{}
	for _, v := range strings.Split(string(f), "\n") {
		line := []rune{}
		if len(v) == 0 {
			continue
		}
		for _, c := range v {
			line = append(line, c)
		}
		grid = append(grid, line)
	}
	return grid
}

func countLights(grid [][]rune) int {
	count := 0
	for _, i := range grid {
		for _, j := range i {
			if j == '#' {
				count += 1
			}
		}
	}
	return count
}

func main() {
	grid := read("input")
	for i := 0; i < 100; i++ {
		grid = iterate(grid, false)
	}
	fmt.Println("Part 1 =", countLights(grid))

	grid = read("input")
	cornerLights(grid)
	for i := 0; i < 100; i++ {
		grid = iterate(grid, true)
	}
	fmt.Println("Part 2 =", countLights(grid))
}
