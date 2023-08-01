package main

import (
	"fmt"
	"log"
	"os"
)

func scan(f []byte, lim int) int {
	floor := 0
	for i := range f {
		switch f[i] {
		case '(':
			floor += 1
		case ')':
			floor += -1
		}
		if lim != 0 && floor == lim {
			return i + 1
		}
	}
	return floor
}

func main() {
	f, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1 = ", scan(f, 0))
	fmt.Println("Part 2 = ", scan(f, -1))
}
