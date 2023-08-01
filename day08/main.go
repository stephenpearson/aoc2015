package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(f), "\n")
	p1posCount := 0
	p1byteCount := 0
	p2Count := 0
	for _, v := range lines {
		p := Parser{v, 0, 0}
		e := Encoder{v, 0}
		p.countBytes()
		e.encode()
		p1posCount += p.pos
		p1byteCount += p.byteCount
		fmt.Println(e.bytecount)
		p2Count += e.bytecount - len(v)
	}
	fmt.Println("Part 1 =", p1posCount-p1byteCount)
	fmt.Println("Part 2 =", p2Count)
}
