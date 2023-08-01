package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func read(filename string) []int {
	f, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	sizes := []int{}
	for _, v := range strings.Split(string(f), "\n") {
		n, err := strconv.Atoi(v)
		if err == nil {
			sizes = append(sizes, n)
		}
	}
	return sizes
}

func mapSelections(selections []int, sizes []int) int {
	total := 0
	for _, v := range selections {
		total += sizes[v]
	}
	return total
}

func iterate(selections []int, sizes []int, testSelection func([]int, int, *int) bool) int {
	var start int
	count := 0
	if len(selections) == 0 {
		start = 0
	} else {
		start = selections[len(selections)-1] + 1
	}
	for i := start; i < len(sizes); i++ {
		new := append(selections, i)
		total := mapSelections(new, sizes)
		flag := testSelection(new, total, &count)
		tmp := iterate(new, sizes, testSelection)
		if flag {
			count += tmp
		} else {
			if count == 0 || (tmp > 0 && tmp < count) {
				count = tmp
			}
		}
	}
	return count
}

func main() {
	sizes := read("input")
	count := iterate([]int{}, sizes, func(selections []int, total int, count *int) bool {
		if total == 150 {
			*count += 1
		}
		return true
	})
	fmt.Println("Part 1 =", count)

	min := iterate([]int{}, sizes, func(selections []int, total int, count *int) bool {
		l := len(selections)
		if total == 150 && (*count == 0 || l < *count) {
			*count = l
		}
		return false
	})
	count = iterate([]int{}, sizes, func(selections []int, total int, count *int) bool {
		if total == 150 && len(selections) == min {
			*count += 1
		}
		return true
	})
	fmt.Println("Part 2 =", count)
}
