package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/fatih/color"
)

type grid struct {
	area [][]byte
}

func move(pos, dir [2]int) [2]int {
	return [2]int{pos[0] + dir[0], pos[1] + dir[1]}
}

func rotate(dir [2]int) [2]int {
	return [2]int{dir[1], -dir[0]}
}

func walk(grid [][]byte, pos [2]int, dir [2]int, sleep time.Duration) ([][2]int, bool) {
	visited := make(map[[4]int]struct{})

	visitedToSlice := func() [][2]int {
		unsortedSteps := make([][2]int, 0, len(visited))
		seen := make(map[[2]int]struct{})
		for k := range visited {
			if _, ok := seen[[2]int{k[0], k[1]}]; ok {
				continue
			}
			seen[[2]int{k[0], k[1]}] = struct{}{}
			unsortedSteps = append(unsortedSteps, [2]int{k[0], k[1]})
		}
		return unsortedSteps
	}

OUTER:
	for {
		if sleep > 0 && len(visited) > 3200 {
			print(grid, visitedToSlice(), pos)
			fmt.Printf("%d\n", len(visited))
			time.Sleep(sleep)
		}
		if _, found := visited[[4]int{pos[0], pos[1], dir[0], dir[1]}]; found {
			return visitedToSlice(), true
		}
		visited[[4]int{pos[0], pos[1], dir[0], dir[1]}] = struct{}{}
		for {
			newPos := move(pos, dir)
			if newPos[0] < 0 || newPos[0] == len(grid) || newPos[1] < 0 || newPos[1] == len(grid[1]) {
				break OUTER
			}
			if grid[newPos[0]][newPos[1]] == '#' {
				dir = rotate(dir)
				continue
			}
			pos = newPos
			break
		}
	}

	return visitedToSlice(), false
}

func read(content []byte) ([][]byte, [2]int) {
	grid := make([][]byte, 0)
	startPos := [2]int{}

	s := bufio.NewScanner(bytes.NewReader(content))
	rowIdx := 0
	for s.Scan() {
		line := s.Text()
		grid = append(grid, []byte(line))
		if colIdx := strings.Index(line, "^"); colIdx != -1 {
			startPos = [2]int{rowIdx, colIdx}
		}
		rowIdx++
	}
	return grid, startPos
}

func print(grid [][]byte, visited [][2]int, pos [2]int) {
	var b strings.Builder
	for rowIdx, row := range grid {
		for colIdx, c := range row {
			if rowIdx == obstacle[0] && colIdx == obstacle[1] {
				color.New(color.FgRed).Fprint(&b, "O")
				continue
			}
			if rowIdx == pos[0] && colIdx == pos[1] {
				color.New(color.FgRed).Fprint(&b, "X")
				continue
			}
			if slices.Contains(visited, [2]int{rowIdx, colIdx}) {
				color.New(color.FgYellow).Fprint(&b, ".")
				//fmt.Fprint(&b, ".")
				continue
			}
			if c == '.' {
				c = ' '
			}
			fmt.Fprintf(&b, "%c", c)
		}
		fmt.Fprint(&b, "\n")
	}
	fmt.Printf("%s\n", b.String())
}

var obstacle [2]int

func main() {
	content, _ := os.ReadFile("input.txt")
	grid, start := read(content)
	north := [2]int{-1, 0}
	visited, _ := walk(grid, start, north, 0)
	fmt.Printf("Part A: %d\n", len(visited))

	found := 0
	for _, pos := range visited {
		if grid[pos[0]][pos[1]] != '.' {
			continue
		}
		grid[pos[0]][pos[1]] = '#'
		obstacle = pos // for printing

		if _, loop := walk(grid, start, north, 0); loop {
			found++
		}
		grid[pos[0]][pos[1]] = '.'
	}
	fmt.Printf("Part B: %d\n", found)
}
