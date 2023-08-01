package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type gate struct {
	input_a string
	input_b string
	output  string
	fn      func(uint16, uint16) uint16
}

func gate_set(a uint16, b uint16) uint16 {
	return a
}

func gate_not(a uint16, b uint16) uint16 {
	return a ^ 0xFFFF
}

func gate_and(a uint16, b uint16) uint16 {
	return a & b
}

func gate_or(a uint16, b uint16) uint16 {
	return a | b
}

func gate_rshift(a uint16, b uint16) uint16 {
	return a >> b
}

func gate_lshift(a uint16, b uint16) uint16 {
	return a << b
}

func iterate(gates *[]gate, wires *map[string]uint16) {
	for i := 0; i < len(*gates); i++ {

		var input_a uint16
		a, e := strconv.Atoi((*gates)[i].input_a)
		if e == nil {
			input_a = uint16(a)
		} else {
			a, ok := (*wires)[(*gates)[i].input_a]
			if ok {
				input_a = a
			} else {
				continue
			}
		}

		var input_b uint16
		b, e := strconv.Atoi((*gates)[i].input_b)
		if e == nil {
			input_b = uint16(b)
		} else {
			b, ok := (*wires)[(*gates)[i].input_b]
			if ok {
				input_b = b
			} else {
				continue
			}
		}

		(*wires)[(*gates)[i].output] = (*gates)[i].fn(input_a, input_b)
	}
}

func main() {
	f, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(f), "\n")
	wires := make(map[string]uint16)
	gates := make([]gate, 0)
	for i := range lines {
		fields := strings.Split(lines[i], " ")
		if len(fields) <= 1 {
			continue
		} else if len(fields) == 3 {
			gates = append(gates, gate{input_a: fields[0], input_b: "0", output: fields[2], fn: gate_set})
		} else if fields[0] == "NOT" {
			gates = append(gates, gate{input_a: fields[1], input_b: "0", output: fields[3], fn: gate_not})
		} else if fields[1] == "AND" {
			gates = append(gates, gate{input_a: fields[0], input_b: fields[2], output: fields[4], fn: gate_and})
		} else if fields[1] == "OR" {
			gates = append(gates, gate{input_a: fields[0], input_b: fields[2], output: fields[4], fn: gate_or})
		} else if fields[1] == "RSHIFT" {
			gates = append(gates, gate{input_a: fields[0], input_b: fields[2], output: fields[4], fn: gate_rshift})
		} else if fields[1] == "LSHIFT" {
			gates = append(gates, gate{input_a: fields[0], input_b: fields[2], output: fields[4], fn: gate_lshift})
		}
	}
	var v1 uint16
	var ok bool
	for {
		iterate(&gates, &wires)
		v1, ok = wires["a"]
		if ok {
			fmt.Println("Part 1 =", v1)
			break
		}
	}

	for i := 0; i < len(gates); i++ {
		if gates[i].output == "b" {
			gates[i].input_a = fmt.Sprintf("%d", v1)
		}
	}

	wires = make(map[string]uint16)
	var v2 uint16
	for {
		iterate(&gates, &wires)
		v2, ok = wires["a"]
		if ok {
			fmt.Println("Part 2 =", v2)
			break
		}
	}
}
