package main

import (
	"fmt"
)

func rule1(str []rune) bool {
	for i := 0; i < len(str)-2; i++ {
		if str[i] == str[i+1]-1 && str[i] == str[i+2]-2 {
			return true
		}
	}
	return false
}

func rule2(str []rune) bool {
	for _, v := range str {
		if v == 'i' || v == 'o' || v == 'l' {
			return false
		}
	}
	return true
}

func rule3(str []rune) bool {
	for i := 0; i < len(str)-3; i++ {
		if str[i] == str[i+1] {
			for j := i + 2; j < len(str)-1; j++ {
				if str[j] == str[j+1] {
					return true
				}
			}
		}
	}
	return false
}

func nextPassword(str []rune) ([]rune, bool) {
	p := len(str) - 1
	result := []rune(str)
	more := true
	for {
		next := rune(result[p] + 1)
		if next > 'z' {
			result[p] = 'a'
			p += -1
			if p < 0 {
				more = false
			}
		} else {
			result[p] = next
			break
		}
	}
	return result, more
}

func main() {
	password := "cqjxjnds"
	str := []rune(password)
	var more bool
	count := 0
	for count < 2 {
		str, more = nextPassword(str)
		if !more {
			break
		}
		if rule1(str) && rule2(str) && rule3(str) {
			fmt.Println("Part", count+1, "=", string(str))
			count += 1
		}
	}
}
