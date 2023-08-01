package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func find(prev *[]int, val int) bool {
	for _, v := range *prev {
		if v == val {
			return true
		}
	}
	return false
}

func calcDistance(prev *[]int, places *[]string, distances *map[fromTo]int) int {
	total := 0
	for i := 0; i < len(*prev)-1; i++ {
		from := (*places)[(*prev)[i]]
		to := (*places)[(*prev)[i+1]]
		ft := fromTo{from, to}
		dist, ok := (*distances)[ft]
		if ok {
			total += dist
		} else {
			log.Fatal(ft, "unknown")
		}
	}
	return total
}

func iterate(prev []int, places *[]string, distances *map[fromTo]int, longMode bool) int {
	if len(prev) == len(*places) {
		return calcDistance(&prev, places, distances)
	}
	result := 0
	for i := range *places {
		if find(&prev, i) {
			continue
		}
		dist := iterate(append(prev, i), places, distances, longMode)
		if result == 0 || (longMode && dist > result) || (!longMode && dist < result) {
			result = dist
		}
	}
	return result
}

type fromTo struct {
	from string
	to   string
}

func main() {
	f, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	tmpPlaces := make(map[string]bool, 0)
	distances := make(map[fromTo]int, 0)
	for _, line := range strings.Split(string(f), "\n") {
		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}
		tmpPlaces[fields[0]] = true
		tmpPlaces[fields[2]] = true
		dist, _ := strconv.Atoi(fields[4])
		distances[fromTo{fields[0], fields[2]}] = dist
		distances[fromTo{fields[2], fields[0]}] = dist
	}
	var places []string
	for i := range tmpPlaces {
		places = append(places, i)
	}
	fmt.Println("Part 1 =", iterate([]int{}, &places, &distances, false))
	fmt.Println("Part 2 =", iterate([]int{}, &places, &distances, true))
}
