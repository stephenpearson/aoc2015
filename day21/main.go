package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseInts(fields []string) [3]int {
	result := [3]int{}
	for i := 0; i < 3; i++ {
		result[i], _ = strconv.Atoi(fields[i])
	}
	return result
}

type itemset = map[string]map[string][3]int

func readItems(filename string) itemset {
	f, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	var key string
	items := itemset{}
	for _, l := range strings.Split(string(f), "\n") {
		fields := strings.Fields(l)
		if len(fields) == 0 {
			continue
		}
		f0 := strings.Split(fields[0], ":")
		if len(f0) > 1 {
			key = f0[0]
			if _, ok := items[key]; !ok {
				items[key] = map[string][3]int{}
			}
		} else if len(fields) == 4 {
			items[key][fields[0]] = parseInts(fields[1:4])
		} else {
			items[key][fields[0]+" "+fields[1]] = parseInts(fields[2:5])
		}
	}
	return items
}

func iter_weapons(ch chan [3]int, items itemset, cost int, damage int, armor int) {
	for _, v := range items["Weapons"] {
		iter_armor(ch, items, cost+v[0], damage+v[1], armor+v[2])
	}
	close(ch)
}

func iter_armor(ch chan [3]int, items itemset, cost int, damage int, armor int) {
	iter_rings(ch, items, cost, damage, armor)
	for _, v := range items["Armor"] {
		iter_rings(ch, items, cost+v[0], damage+v[1], armor+v[2])
	}
}

func iter_rings(ch chan [3]int, items itemset, cost int, damage int, armor int) {
	ch <- [3]int{cost, damage, armor}

	for _, v2 := range items["Rings"] {
		ch <- [3]int{cost + v2[0], damage + v2[1], armor + v2[2]}
	}

	for i1, v1 := range items["Rings"] {
		ch <- [3]int{cost + v1[0], damage + v1[1], armor + v1[2]}
		for i2, v2 := range items["Rings"] {
			if i2 == i1 {
				continue
			}
			ch <- [3]int{cost + v1[0] + v2[0], damage + v1[1] + v2[1], armor + v1[2] + v2[2]}
		}
	}
}

type Combatant struct {
	hp     int
	damage int
	armor  int
}

func fight(player Combatant, boss Combatant) bool {
	for {
		attack := player.damage - boss.armor
		if attack < 0 {
			attack = 1
		}
		boss.hp -= attack
		if boss.hp <= 0 {
			return true
		}

		attack = boss.damage - player.armor
		if attack < 0 {
			attack = 1
		}
		player.hp -= attack
		if player.hp <= 0 {
			return false
		}
	}
}

func main() {
	items := readItems("items.txt")
	ch := make(chan [3]int)

	go iter_weapons(ch, items, 0, 0, 0)

	boss := Combatant{104, 8, 1}

	lowest, highest := 0, 0
	for i := range ch {
		cost, dam, arm := i[0], i[1], i[2]
		player := Combatant{100, dam, arm}
		if fight(player, boss) {
			if (cost < lowest) || lowest == 0 {
				lowest = cost
			}
		} else {
			if cost > highest {
				highest = cost
			}
		}
	}
	fmt.Println("Part 1 =", lowest)
	fmt.Println("Part 2 =", highest)
}
