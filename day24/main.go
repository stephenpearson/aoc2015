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

	list := []int{}
	for _, i := range strings.Split(string(f), "\n") {
		v, e := strconv.Atoi(i)
		if e == nil {
			list = append(list, v)
		}
	}
	return list
}

func contains(sl []int, val int) bool {
	for i := 0; i < len(sl) && sl[i] <= val; i++ {
		if val == sl[i] {
			return true
		}
	}
	return false
}

func pick(ch chan []int, list []int, qty int, start int, acc []int, excl []int, target int) bool {
	if sum(acc) > target {
		return false
	}
	if qty == 0 && (target == sum(acc)) {
		ch <- acc
		return true
	}
	for i := start; i < len(list); i++ {
		v := list[i]
		if contains(excl, v) {
			continue
		}
		if pick(ch, list, qty-1, i+1, append(acc, list[i]), excl, target) {
			return true
		}
	}
	return false
}

func group1(ch1 chan []int, list []int, target int) {
	for i := 1; i < len(list)-1; i++ {
		if pick(ch1, list, i, 0, []int{}, []int{}, target) {
			break
		}
	}
	close(ch1)
}

func sum(list []int) int {
	s := 0
	for _, v := range list {
		s += v
	}
	return s
}

func iterate(list []int, target int) int {
	ch1 := make(chan []int)
	go group1(ch1, list, target)
	product := 1
	for g1 := range ch1 {
		for _, v := range g1 {
			product = product * v
		}
	}
	return product
}

func main() {
	list := read("input")
	total := 0
	for _, v := range list {
		total += v
	}
	fmt.Println("Part 1 =", iterate(list, total/3))
	fmt.Println("Part 2 =", iterate(list, total/4))
}
