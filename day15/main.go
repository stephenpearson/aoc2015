package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func intVal(str string) int {
	result, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal("Could not parse:", err)
	}
	return result
}

func parseRows(f []byte) [][]int {
	columns := []int{2, 4, 6, 8, 10}
	rows := [][]int{}
	for _, line := range strings.Split(string(f), "\n") {
		row := []int{}
		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}
		for _, i := range columns {
			col := intVal(strings.TrimSuffix(fields[i], ","))
			row = append(row, col)
		}
		rows = append(rows, row)
	}
	return rows
}

func calories(qty []int, ing [][]int) int {
	i := len(ing[0]) - 1
	cals := 0
	for j := 0; j < len(qty); j++ {
		cals += qty[j] * ing[j][i]
	}
	return cals
}

func calc(qty []int, ing [][]int) int {
	score := 1
	for i := 0; i < len(ing[0])-1; i++ {
		prop := 0
		for j := 0; j < len(qty); j++ {
			prop += qty[j] * ing[j][i]
		}
		if prop < 0 {
			prop = 0
		}
		score *= prop
	}
	return score
}

func iterate(prev []int, teaspoons int, ing [][]int, calLimit int) int {
	if teaspoons > 100 {
		return 0
	}
	if len(prev) == len(ing)-1 {
		prev = append(prev, 100-teaspoons)
		if calLimit == 0 || calories(prev, ing) == calLimit {
			return calc(prev, ing)
		} else {
			return 0
		}
	}

	max := 0
	for i := 0; i < 100; i++ {
		val := iterate(append(prev, i), teaspoons+i, ing, calLimit)
		if val > max {
			max = val
		}
	}
	return max
}

func main() {
	f, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	ing := parseRows(f)
	fmt.Println("Part 1 =", iterate([]int{}, 0, ing, 0))
	fmt.Println("Part 2 =", iterate([]int{}, 0, ing, 500))
}
