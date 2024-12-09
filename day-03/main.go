package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type mulInstr string

func (m mulInstr) Eval() int {
	n := m[len("mul(") : len(m)-1]
	n2 := strings.Split(string(n), ",")
	n1i, _ := strconv.Atoi(n2[0])
	n2i, _ := strconv.Atoi(n2[1])
	return n1i * n2i
}

type mulmap struct {
	enableMap  map[int]bool
	enableList []int
}

func (m *mulmap) AddInstruction(pos int, isEnabled bool) {
	m.enableMap[pos] = isEnabled
	m.enableList = append(m.enableList, pos)
	slices.Sort(m.enableList)
}

func (m *mulmap) Enabled(position int) bool {
	lastPos := 0
	for _, pos := range m.enableList {
		if pos < position {
			lastPos = pos
			continue
		}
		break
	}
	if lastPos == 0 {
		return true
	}
	return m.enableMap[lastPos]
}

func part2(content []byte) int {
	mm := &mulmap{
		enableMap:  make(map[int]bool),
		enableList: make([]int, 0),
	}

	instructions := []string{"do()", "don't()"}

	for _, instr := range instructions {
		re := regexp.MustCompile(instr)
		idx := re.FindAllIndex(content, -1)
		for _, i := range idx {
			mm.AddInstruction(i[0], instr == "do()")
		}
	}

	sum := 0
	mulIdx := regexp.MustCompile(`mul\(\d+,\d+\)`).FindAllIndex(content, -1)
	for _, i := range mulIdx {
		if mm.Enabled(i[0]) {
			sum += mulInstr(string(content[i[0]:i[1]])).Eval()
		}
	}

	return sum
}

func main() {
	content, _ := os.ReadFile("input.txt")

	mul := regexp.MustCompile(`mul\(\d+,\d+\)`)
	muls := mul.FindAllString(string(content), -1)
	sum := 0
	for _, m := range muls {
		mInstr := mulInstr(m)
		sum += mInstr.Eval()
	}
	fmt.Printf("Part I: %d\n", sum)
	fmt.Printf("Part II: %d\n", part2(content))

}
