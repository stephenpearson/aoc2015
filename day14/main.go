package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type reindeer struct {
	name      string
	flySpeed  int
	flyTime   int
	restTime  int
	distance  int
	flyClock  int
	restClock int
	points    int
}

func (p *reindeer) move() {
	if p.restClock == 0 {
		if p.flyClock == 0 {
			p.flyClock = p.flyTime
			p.restClock = p.restTime
		}
	}
	if p.flyClock > 0 {
		p.distance += p.flySpeed
		p.flyClock += -1
	} else {
		p.restClock += -1
	}
}

func winner(list []*reindeer) int {
	max := 0
	result := 0
	for i, v := range list {
		if v.distance > max {
			max = v.distance
			result = i
		}
	}
	list[result].points += 1
	return result
}

func main() {
	f, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	list := []*reindeer{}
	for _, line := range strings.Split(string(f), "\n") {
		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}
		flySpeed, _ := strconv.Atoi(fields[3])
		flyTime, _ := strconv.Atoi(fields[6])
		restTime, _ := strconv.Atoi(fields[13])
		r := reindeer{fields[0], flySpeed, flyTime, restTime, 0, 0, 0, 0}
		list = append(list, &r)
	}

	var w int
	for i := 0; i < 2503; i++ {
		for _, v := range list {
			v.move()
		}
		w = winner(list)
	}
	fmt.Println("Part 1:", list[w].name, "=", list[w].distance)

	max := 0
	w = 0
	for i, v := range list {
		if v.points > max {
			w = i
			max = v.points
		}
	}
	fmt.Println("Part 2:", list[w].name, "=", list[w].points)
}
