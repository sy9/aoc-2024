package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func readInput(filename string) ([]int, []int) {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic("unable to read file: " + err.Error())
	}
	left := make([]int, 0)
	right := make([]int, 0)
	for _, line := range strings.Split(string(content), "\n") {
		segments := strings.Fields(line)
		i, _ := strconv.Atoi(segments[0])
		left = append(left, i)
		i, _ = strconv.Atoi(segments[1])
		right = append(right, i)
	}
	return left, right
}

func main() {
	left, right := readInput("input.txt")
	slices.Sort(left)
	slices.Sort(right)
	counts := make(map[int]int)
	for i := range right {
		counts[right[i]]++
	}
	sum := 0
	sim := 0 // similarity
	for i := range left {
		d := right[i] - left[i]
		if d < 0 {
			d *= -1
		}
		sum += d
		c, ok := counts[left[i]]
		if !ok {
			continue
		}
		sim += left[i] * c
	}
	fmt.Printf("Part I: %d\n", sum)
	fmt.Printf("Part II: %d\n", sim)
}
