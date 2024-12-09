package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

// line
func line(a, b [2]int, w, h int) [][2]int {
	positions := make([][2]int, 0)

	rowOffset := a[0] - b[0]
	colOffset := a[1] - b[1]

	rowPos := a[0]
	colPos := a[1]
	for rowPos < h && colPos < w && rowPos >= 0 && colPos >= 0 {
		positions = append(positions, [2]int{rowPos, colPos})
		rowPos += rowOffset
		colPos += colOffset
	}

	rowPos = a[0] - rowOffset
	colPos = a[1] - colOffset

	for rowPos < h && colPos < w && rowPos >= 0 && colPos >= 0 {
		positions = append(positions, [2]int{rowPos, colPos})
		rowPos -= rowOffset
		colPos -= colOffset
	}

	return positions
}

// reflect reflects around a
func reflect(a, b [2]int, w, h int) [][2]int {
	row := a[0] + (a[0] - b[0])
	col := a[1] + (a[1] - b[1])

	if row < 0 || col < 0 || row >= h || col >= w {
		return nil
	}
	return [][2]int{{row, col}}
}

func analize(positions [][2]int, w, h int, find func(a, b [2]int, w, h int) [][2]int) [][2]int {
	antinodes := make([][2]int, 0)

	for _, a := range positions {
		for _, b := range positions {
			if a == b {
				continue
			}
			pos := find(a, b, w, h)
			antinodes = append(antinodes, pos...)
			pos = find(b, a, w, h)
			antinodes = append(antinodes, pos...)
		}
	}
	return antinodes
}

func partA(m map[string][][2]int, w, h int) int {
	uniquePositions := make(map[[2]int]struct{})
	for _, positions := range m {
		antinodes := analize(positions, w, h, reflect)
		for _, an := range antinodes {
			uniquePositions[an] = struct{}{}
		}
	}
	return len(uniquePositions)
}

func partB(m map[string][][2]int, w, h int) int {
	uniquePositions := make(map[[2]int]struct{})
	for _, positions := range m {
		antinodes := analize(positions, w, h, line)
		for _, an := range antinodes {
			uniquePositions[an] = struct{}{}
		}
	}
	return len(uniquePositions)
}
func parse(content []byte) (m map[string][][2]int, w, h int) {
	s := bufio.NewScanner(bytes.NewReader(content))
	m = make(map[string][][2]int)
	for s.Scan() {
		row := s.Text()
		w = len(row)
		for col, content := range row {
			if content == '.' {
				continue
			}
			m[string(content)] = append(m[string(content)], [2]int{h, col})
		}
		h++
	}
	return
}

func main() {
	content, _ := os.ReadFile("input.txt")
	m, w, h := parse(content)
	fmt.Printf("Part A: %d\n", partA(m, w, h))
	fmt.Printf("Part B: %d\n", partB(m, w, h))
}
