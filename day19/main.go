package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type rule struct {
	from string
	to   string
}

func substr(str string, sub string, pos int) bool {
	for i := range sub {
		if pos+i >= len(str) || str[pos+i] != sub[i] {
			return false
		}
	}
	return true
}

func contains(str string, sub string) bool {
	for i := 0; i <= len(str)-len(sub); i++ {
		if str[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}

func generator(ch chan string, rules []rule, input string) {
	if contains(input, "e") && len(input) > 1 {
		close(ch)
		return
	}
	for _, r := range rules {
		for p := 0; p < len(input); {
			if substr(input, r.from, p) {
				output := input[:p] + r.to + input[p+len(r.from):]
				ch <- output
				p += len(r.from)
			} else {
				p += 1
			}
		}
	}
	close(ch)
}

func read(filename string) ([]rule, string) {
	f, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	rules := []rule{}
	var input string
	for _, v := range strings.Split(string(f), "\n") {
		fields := strings.Fields(v)
		if len(fields) == 3 {
			rules = append(rules, rule{fields[0], fields[2]})
		} else if len(fields) == 1 {
			input = fields[0]
		}
	}
	return rules, input
}

func iterate(input string, rules []rule) int {
	store := map[string]int{}
	tried := map[string]bool{}
	store[input] = 0
	var depth int

	count := 0
	for {
		shortest := ""
		for i, v := range store {
			if i == "e" {
				return v
			}
			if !tried[i] && len(i) < len(shortest) || shortest == "" {
				shortest = i
				depth = v + 1
			}
		}
		ch := make(chan string)
		go generator(ch, rules, shortest)
		for str := range ch {
			store[str] = depth
		}
		tried[shortest] = true
		delete(store, shortest)
		count++
		if count > 300 { // bodged it
			return 0
		}
	}
}

func retry(input string, rules []rule) int {
	for {
		result := iterate(input, rules)
		if result != 0 {
			return result
		}
	}
}

func main() {
	ch := make(chan string)
	rules, input := read("input")

	result := map[string]bool{}
	go generator(ch, rules, input)
	for i := range ch {
		result[i] = true
	}
	fmt.Println("Part 1 =", len(result))

	reverse_rules := []rule{}
	for _, v := range rules {
		reverse_rules = append(reverse_rules, rule{v.to, v.from})
	}
	fmt.Println("Part 2 =", retry(input, reverse_rules))
}
