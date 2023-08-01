package main

import (
	"fmt"
)

func lookSay(str []int) []int {
	cur := 0
	count := 0
	acc := make([]int, 0)
	for i := 0; i <= len(str); i++ {
		var val int
		if i < len(str) {
			val = str[i]
		} else {
			val = 0
		}
		if cur == 0 {
			cur = val
			count += 1
		} else if cur == val {
			count += 1
		} else {
			acc = append(acc, count)
			acc = append(acc, cur)
			count = 1
			cur = val
		}
	}
	return acc
}

func main() {
	str := []int{1, 3, 2, 1, 1, 3, 1, 1, 1, 2}
	for i := 0; i < 40; i++ {
		str = lookSay(str)
	}
	fmt.Println("Part 1 =", len(str))
	for i := 0; i < 10; i++ {
		str = lookSay(str)
	}
	fmt.Println("Part 2 =", len(str))
}
