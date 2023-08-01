package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

func part1(input string, f func(r [16]byte) bool) int {
	number := 0
	for {
		data := []byte(input + strconv.Itoa(number))
		result := md5.Sum(data)
		if f(result) {
			return number
		}
		number += 1
	}
}

func main() {
	input := "ckczppom"
	fmt.Println("Part 1 =", part1(input, func(r [16]byte) bool {
		return r[0]|r[1]|(r[2]&0xF0) == 0
	}))
	fmt.Println("Part 2 =", part1(input, func(r [16]byte) bool {
		return r[0]|r[1]|r[2] == 0
	}))
}
