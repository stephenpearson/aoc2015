package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	opcode   string
	operands []string
}

func read(filename string) []Instruction {
	f, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	code := []Instruction{}
	for _, v := range strings.Split(string(f), "\n") {
		fields := strings.Fields(strings.Replace(v, ",", "", 1))
		if len(fields) == 0 {
			continue
		}
		code = append(code, Instruction{fields[0], fields[1:]})
	}
	return code
}

type Registers struct {
	reg_pc uint
	reg_a  uint
	reg_b  uint
}

func apply(rname string, regs *Registers, param int, fn func(*Registers, uint) uint) {
	if rname == "a" {
		regs.reg_a = fn(regs, regs.reg_a)
	} else if rname == "b" {
		regs.reg_b = fn(regs, regs.reg_b)
	} else if rname == "pc" {
		regs.reg_pc = fn(regs, regs.reg_pc)
	}
}

func execute(code []Instruction, regs *Registers) {
	operands := code[regs.reg_pc].operands
	switch code[regs.reg_pc].opcode {
	case "hlf":
		apply(operands[0], regs, 0, func(r *Registers, i uint) uint {
			r.reg_pc += 1
			return i >> 1
		})
	case "tpl":
		apply(operands[0], regs, 0, func(r *Registers, i uint) uint {
			r.reg_pc += 1
			return i * 3
		})
	case "inc":
		apply(operands[0], regs, 0, func(r *Registers, i uint) uint {
			r.reg_pc += 1
			return i + 1
		})
	case "jmp":
		offset, _ := strconv.Atoi(operands[0])
		apply("pc", regs, offset, func(r *Registers, i uint) uint {
			return uint(int(i) + offset)
		})
	case "jie":
		offset, _ := strconv.Atoi(operands[1])
		apply(operands[0], regs, offset, func(r *Registers, i uint) uint {
			if i%2 == 0 {
				r.reg_pc = uint(int(r.reg_pc) + offset)
			} else {
				r.reg_pc += 1
			}
			return i
		})
	case "jio":
		offset, _ := strconv.Atoi(operands[1])
		apply(operands[0], regs, offset, func(r *Registers, i uint) uint {
			if i == 1 {
				r.reg_pc = uint(int(r.reg_pc) + offset)
			} else {
				r.reg_pc += 1
			}
			return i
		})
	}
}

func run(code []Instruction, regs *Registers) uint {
	for int(regs.reg_pc) < len(code) {
		execute(code, regs)
	}
	return regs.reg_b
}

func main() {
	regs := Registers{}
	code := read("input")
	fmt.Println("Part 1 =", run(code, &regs))

	regs = Registers{}
	regs.reg_a = 1
	fmt.Println("Part 2 =", run(code, &regs))
}
