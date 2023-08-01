package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func filter1(aunts []map[string]int, key string, value int) []map[string]int {
	result := []map[string]int{}
	for _, v := range aunts {
		val, ok := v[key]
		if !ok || val == value {
			result = append(result, v)
		}
	}
	return result
}

func filter2(aunts []map[string]int, key string, value int) []map[string]int {
	result := []map[string]int{}
	for _, v := range aunts {
		val, ok := v[key]
		if !ok {
			result = append(result, v)
		} else if key == "cats" || key == "trees" {
			if val > value {
				result = append(result, v)
			}
		} else if key == "pomeranians" || key == "goldfish" {
			if val < value {
				result = append(result, v)
			}
		} else if val == value {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	f, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	aunts1 := []map[string]int{}
	for i, line := range strings.Split(string(f), "\n") {
		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}
		things := map[string]int{"aunt": i + 1}
		for i := 2; i < len(fields); i += 2 {
			key := strings.TrimSuffix(fields[i], ":")
			value, err := strconv.Atoi(strings.TrimSuffix(fields[i+1], ","))
			if err != nil {
				log.Fatal("Could not parse:", fields[i+1])
			}
			things[key] = value
		}
		aunts1 = append(aunts1, things)
	}

	aunts2 := make([]map[string]int, len(aunts1))
	copy(aunts2, aunts1)

	exclusions := map[string]int{"children": 3, "cats": 7, "samoyeds": 2, "pomeranians": 3, "akitas": 0, "vizslas": 0, "goldfish": 5, "trees": 3, "cars": 2, "perfumes": 1}
	for k, v := range exclusions {
		aunts1 = filter1(aunts1, k, v)
	}
	fmt.Println("Part 1 =", aunts1[0]["aunt"])

	for k, v := range exclusions {
		aunts2 = filter2(aunts2, k, v)
	}
	fmt.Println("Part 2 =", aunts2[0]["aunt"])
}
