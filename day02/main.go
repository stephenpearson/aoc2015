package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func min2(v1 int, v2 int, v3 int) (int, int) {
	list := []int{v1, v2, v3}
	sort.Slice(list, func(i, j int) bool {
		return list[i] < list[j]
	})
	return list[0], list[1]
}

func min3(v1 int, v2 int, v3 int) int {
	result := v1
	if v2 < result {
		result = v2
	}
	if v3 < result {
		result = v3
	}
	return result
}

func main() {
	f, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(f), "\n")
	total := 0
	ribbon := 0
	for i := range lines {
		vals := strings.Split(lines[i], "x")
		if len(vals) > 1 {
			l, _ := strconv.Atoi(vals[0])
			w, _ := strconv.Atoi(vals[1])
			h, _ := strconv.Atoi(vals[2])
			lw, wh, lh := l*w, w*h, l*h
			m := min3(lw, wh, lh)
			area := 2*lw + 2*wh + 2*lh + m
			total += area
			a, b := min2(l, w, h)
			ribbon += a + a + b + b + l*w*h
		}
	}
	fmt.Println("Part 1 =", total)
	fmt.Println("Part 2 =", ribbon)
}
