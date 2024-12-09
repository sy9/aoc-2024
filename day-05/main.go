package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type printingRules struct {
	after         map[int]map[int]struct{} // after includes all page numbers that must occur after the given page
	notValidIndex int                      // Valid() sets this
}

func (pr *printingRules) Add(page, after int) {
	pages, ok := pr.after[page]
	if !ok {
		pr.after[page] = map[int]struct{}{after: {}}
		return
	}
	pages[after] = struct{}{}
}

// Valid checks if pages listed in before do not violate rules for given page.
func (pr *printingRules) Valid(page int, before []int) bool {
	pages, ok := pr.after[page]
	if !ok {
		return true // no rules found for given page
	}

	for i, beforePage := range before {
		if _, found := pages[beforePage]; found {
			pr.notValidIndex = i
			return false // given page should be after, but is before
		}
	}
	return true
}

// Sort sorts update according to the printing rules and returns
// the value of the middle item.
func (pr *printingRules) Sort(update []int) int {
OUTER:
	for {
		for i, page := range update {
			if !pr.Valid(page, update[:i]) {
				update[pr.notValidIndex], update[i] = update[i], update[pr.notValidIndex]
				continue OUTER
			}
		}
		break
	}
	return update[len(update)/2]
}

func read(filename string) (*printingRules, [][]int) {
	pr := &printingRules{
		after: make(map[int]map[int]struct{}),
	}
	updates := make([][]int, 0)
	content, _ := os.ReadFile(filename)
	instructionMode := true
	for _, line := range strings.Split(string(content), "\n") {
		if len(line) == 0 {
			instructionMode = false
			continue
		}

		if instructionMode {
			nums := strings.Split(line, "|")
			a, _ := strconv.Atoi(nums[0])
			b, _ := strconv.Atoi(nums[1])
			pr.Add(a, b)
			continue
		}

		nums := strings.Split(line, ",")
		update := make([]int, 0, len(nums))
		for _, num := range nums {
			numI, _ := strconv.Atoi(num)
			update = append(update, numI)
		}
		updates = append(updates, update)
	}
	return pr, updates
}

func main() {
	pr, updates := read("input.txt")
	partI := 0
	partII := 0

OUTER:
	for _, update := range updates {
		for i, page := range update {
			if !pr.Valid(page, update[:i]) {
				partII += pr.Sort(update)
				continue OUTER
			}
		}
		partI += update[len(update)/2]
	}

	fmt.Printf("Part I: %d\n", partI)
	fmt.Printf("Part II: %d\n", partII)
}
