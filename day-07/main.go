package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func addOp(a, b int) int {
	return a + b
}

func mulOp(a, b int) int {
	return a * b
}

func concatOp(a, b int) int {
	concat, _ := strconv.Atoi(fmt.Sprintf("%d%d", a, b))
	return concat
}

type equation struct {
	testValue int
	numbers   []int
	ops       []func(a, b int) int
}

func newEquation(testValue int, numbers []int) *equation {
	return &equation{
		testValue: testValue,
		numbers:   numbers,
	}
}

func (e *equation) addOperator(op func(a, b int) int) *equation {
	e.ops = append(e.ops, op)
	return e
}

// 292: 11 6 16 20
func (e *equation) calibrate() bool {
	return e.calc(e.testValue, e.numbers[0], 1)
}

func (e *equation) calc(target, total int, idx int) bool {
	if target-total == 0 && idx == len(e.numbers) {
		return true
	}
	if target-total < 0 || idx == len(e.numbers) {
		return false
	}
	for _, op := range e.ops {
		if e.calc(target, op(total, e.numbers[idx]), idx+1) {
			return true
		}
	}
	return false
}

func read(content []byte) []*equation {
	s := bufio.NewScanner(bytes.NewReader(content))
	equations := make([]*equation, 0)
	for s.Scan() {
		parts := strings.Split(s.Text(), ": ")
		testValue, _ := strconv.Atoi(parts[0])
		numsStr := strings.Fields(parts[1])
		nums := make([]int, 0, len(numsStr))
		for _, numStr := range numsStr {
			num, _ := strconv.Atoi(numStr)
			nums = append(nums, num)
		}
		equations = append(equations, newEquation(testValue, nums).addOperator(addOp).addOperator(mulOp))
	}
	return equations
}
func main() {
	f, _ := os.ReadFile("input.txt")
	equations := read(f)
	sum := 0
	for _, eq := range equations {
		if eq.calibrate() {
			sum += eq.testValue
		}
	}
	fmt.Printf("Part A: %d\n", sum)

	sum = 0
	for _, eq := range equations {
		eq.addOperator(concatOp)
		if eq.calibrate() {
			sum += eq.testValue
		}
	}
	fmt.Printf("Part B: %d\n", sum)
}
