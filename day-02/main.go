package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func readFile(filename string) [][]int {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic("unable to read file: " + err.Error())
	}
	reports := make([][]int, 0)
	for _, report := range strings.Split(string(content), "\n") {
		samples := strings.Fields(report)
		reportI := make([]int, 0, len(samples))
		for i := range samples {
			sampleI, _ := strconv.Atoi(samples[i])
			reportI = append(reportI, sampleI)
		}
		reports = append(reports, reportI)
	}
	return reports
}

func safetyDetector(r []int, threshold int, problemDampener bool) bool {
	isSave, offset := scanReport(r, threshold)
	if !problemDampener {
		return isSave
	}
	if !isSave {
		if offset > 0 {
			offset -= 1
		}
		for i := offset; i < len(r) && i < offset+3; i++ {
			rd := slices.Delete(slices.Clone(r), i, i+1)
			if isSave, _ := scanReport(rd, threshold); isSave {
				return isSave
			}
		}
	}
	return isSave
}

func scanReport(r []int, threshold int) (bool, int) {
	direction := 0
	for i := range r[:len(r)-1] {
		if direction == 0 {
			if r[i] < r[i+1] {
				direction = -1 // increase
			} else {
				direction = 1
			}
		}
		diff := (r[i] - r[i+1]) * direction
		if diff > threshold || diff <= 0 {
			return false, i
		}
	}
	return true, 0
}

func main() {
	reports := readFile("input.txt")
	safeCount, safeCount2 := 0, 0
	for _, r := range reports {
		s1 := safetyDetector(r, 3, false)
		s2 := safetyDetector(r, 3, true)
		if s1 {
			safeCount++
		}
		if s2 {
			safeCount2++
		}
	}
	fmt.Printf("Part I: %d\n", safeCount)
	fmt.Printf("Part II: %d\n", safeCount2)
}
