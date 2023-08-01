package main

import "fmt"

func main() {
	row, col := 1, 1
	var v uint64 = 20151125
	for {
		if row == 3010 && col == 3019 {
			fmt.Println("Part 1", "=", v)
			break
		}
		v = (v * 252533) % 33554393
		if row == 1 {
			row = col + 1
			col = 1
		} else {
			row--
			col++
		}
	}
}
