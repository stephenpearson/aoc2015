package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func findName(names []string, name string) int {
	for i, v := range names {
		if v == name {
			return i
		}
	}
	return -1
}

func getNames(f []byte) []string {
	names := []string{}
	for _, line := range strings.Split(string(f), "\n") {
		line = strings.TrimSuffix(line, ".")
		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}
		if findName(names, fields[0]) == -1 {
			names = append(names, fields[0])
		}
		if findName(names, fields[10]) == -1 {
			names = append(names, fields[10])
		}
	}
	return names
}

type pair struct {
	subj int
	obj  int
}

func getHappiness(f []byte, names []string) map[pair]int {
	result := make(map[pair]int, 0)
	for _, line := range strings.Split(string(f), "\n") {
		line = strings.TrimSuffix(line, ".")
		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}
		p := pair{findName(names, fields[0]), findName(names, fields[10])}
		val, _ := strconv.Atoi(fields[3])
		if fields[2] == "lose" {
			val *= -1
		}
		result[p] = val
	}
	return result
}

func findInt(list *[]int, val int) bool {
	for _, v := range *list {
		if v == val {
			return true
		}
	}
	return false
}

func calcHappiness(perm []int, happiness *map[pair]int) int {
	result := 0
	for i := range perm {
		var left, right int
		if i == 0 {
			left = len(perm) - 1
		} else {
			left = i - 1
		}
		if i == len(perm)-1 {
			right = 0
		} else {
			right = i + 1
		}
		p := pair{perm[i], perm[left]}
		result += (*happiness)[p]
		p = pair{perm[i], perm[right]}
		result += (*happiness)[p]
	}
	return result
}

func permute(prev []int, names *[]string, happiness *map[pair]int) int {
	if len(prev) == len(*names) {
		h := calcHappiness(prev, happiness)
		return h
	}
	max := 0
	for i := range *names {
		if findInt(&prev, i) {
			continue
		}
		score := permute(append(prev, i), names, happiness)
		if score > max {
			max = score
		}
	}
	return max
}

func main() {
	f, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	names := getNames(f)
	happiness := getHappiness(f, names)

	fmt.Println("Part 1 =", permute([]int{}, &names, &happiness))
	me := len(names)
	names = append(names, "me")
	for i := 0; i < me; i++ {
		p := pair{i, me}
		happiness[p] = 0
		p = pair{me, 0}
		happiness[p] = 0
	}
	fmt.Println("Part 2 =", permute([]int{}, &names, &happiness))
}
