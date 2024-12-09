package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type valve struct {
	name      string
	flowRate  int
	neighbors []*valve
}

type tunnels struct {
	network map[string]*valve
}

func read(name string) *tunnels {
	t := &tunnels{network: make(map[string]*valve)}
	f, _ := os.Open(name)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		name := line[len("Valve ") : len("Valve ")+3]
		rateStr := strings.Split(line[len("Valve XX has flow rate="):], ";")
		neighbors := make([]string, 0)
		if len(rateStr[1]) == 25 { //  tunnel leads to valve XX
			neighbors = []string{rateStr[1][23:]}
		} else {
			neighbors = strings.Split(rateStr[1][23:], ", ")
		}
		flowRate, _ := strconv.Atoi(rateStr[0])
		v, ok := t.network[name]
		if !ok {
			t.network[name] = &valve{
				name:     name,
				flowRate: flowRate,
			}
		} else {
			v.name = name
			v.flowRate = flowRate
		}
		for _, name := range neighbors {
			v, ok := t.network[name]
			if !ok {
				v = &valve{}
				t.network[name] = v
			}

		}
		fmt.Printf("%s %s %d %v\n", name, rateStr[0], len(rateStr[1]), neighbors)
		break
	}
	return t
}

func main() {
	read("input.txt")
}
