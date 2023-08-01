package main

import "fmt"

func calc(target int, presents int, max int) int {
	cells := make([]int, target)
	for i := 1; i < target; i++ {
		count := 0
		for j := i; j < target; j += i {
			cells[j] += presents * i
			count++
			if max != 0 && count > max {
				break
			}
		}
	}
	for i := range cells {
		if cells[i] >= target {
			return i
		}
	}
	return 0
}

func main() {
	target := 33100000
	fmt.Println("Part 1 =", calc(target, 10, 0))
	fmt.Println("Part 2 =", calc(target, 11, 50))
}
