package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func iterate(data *interface{}, sum float64, part2 bool) float64 {
	switch (*data).(type) {
	case []interface{}:
		for _, v := range (*data).([]interface{}) {
			sum += iterate(&v, 0, part2)
		}
	case map[string]interface{}:
		for _, v := range (*data).(map[string]interface{}) {
			if v == "red" && part2 {
				return 0
			}
			sum += iterate(&v, 0, part2)
		}
	case float64:
		val, ok := (*data).(float64)
		if ok {
			sum += val
		}
	}
	return sum
}

func main() {
	f, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	var data interface{}
	json.Unmarshal(f, &data)
	fmt.Println("Part 1 =", iterate(&data, 0.0, false))
	fmt.Println("Part 2 =", iterate(&data, 0.0, true))
}
