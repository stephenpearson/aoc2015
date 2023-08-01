package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func doubleChar(line string) bool {
	for i := 0; i < len(line)-1; i++ {
		if line[i] == line[i+1] {
			return true
		}
	}
	return false
}

func good1(line string) bool {
	var vowels = regexp.MustCompile(`[aeiou].*[aeiou].*[aeiou]`)
	if !vowels.MatchString(line) {
		return false
	}

	if !doubleChar(line) {
		return false
	}

	bad_strings := []string{"ab", "cd", "pq", "xy"}
	for i := range bad_strings {
		var bad = regexp.MustCompile(bad_strings[i])
		if bad.MatchString(line) {
			return false
		}
	}
	return true
}

func rule21(line string) bool {
	for i := 0; i < len(line)-3; i++ {
		for j := i + 2; j < len(line)-1; j++ {
			if line[i] == line[j] && line[i+1] == line[j+1] {
				return true
			}
		}
	}
	return false
}

func rule22(line string) bool {
	for i := 0; i < len(line)-2; i++ {
		if line[i] == line[i+2] {
			return true
		}
	}
	return false
}

func good2(line string) bool {
	if !rule21(line) || !rule22(line) {
		return false
	}
	return true
}

func count(lines []string, f func(line string) bool) int {
	count := 0
	for i := range lines {
		if f(lines[i]) {
			count += 1
		}
	}
	return count
}

func main() {
	f, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(f), "\n")
	fmt.Println("Part 1 =", count(lines, good1))
	fmt.Println("Part 2 =", count(lines, good2))
}
