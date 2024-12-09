package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type pos struct {
	col, line int
}

type direction int

const (
	dirUnknown direction = iota
	dirUp
	dirUpRight
	dirRight
	dirDownRight
	dirDown
	dirDownLeft
	dirLeft
	dirUpLeft
)

var moveMap = map[direction][]int{ // col, line
	dirUp:        {0, -1},
	dirUpRight:   {1, -1},
	dirRight:     {1, 0},
	dirDownRight: {1, 1},
	dirDown:      {0, 1},
	dirDownLeft:  {-1, 1},
	dirLeft:      {-1, 0},
	dirUpLeft:    {-1, -1},
}

type wordMap struct {
	lines     [][]byte
	startPosX []pos
	startPosA []pos
}

func (wm wordMap) String() string {
	var b strings.Builder
	for _, startPos := range wm.startPosX {
		fmt.Fprintf(&b, "(%d,%d) ", startPos.col, startPos.line)
	}
	return b.String()
}

func (wm wordMap) Search() int {
	directions := []direction{
		dirUp,
		dirUpRight,
		dirRight,
		dirDownRight,
		dirDown,
		dirDownLeft,
		dirLeft,
		dirUpLeft,
	}
	findings := 0

	for _, startPos := range wm.startPosX {
		for _, direction := range directions {
			posCol, posLine := startPos.col, startPos.line
			var ok bool
			found := true // assume
			for _, ch := range "MAS" {
				posCol, posLine, ok = wm.move(posCol, posLine, direction)
				if !ok || wm.lines[posLine][posCol] != byte(ch) {
					found = false
					break
				}
			}
			if found {
				findings++
			}
		}
	}
	return findings
}

func (wm *wordMap) SearchXMAS() int {
	findings := 0
OUTER:
	for _, startPos := range wm.startPosA {
		for _, directions := range [][]direction{
			{dirUpLeft, dirDownRight},
			{dirUpRight, dirDownLeft},
		} {
			toFind := "MS"
			for _, direction := range directions {
				posCol, posLine, ok := wm.move(startPos.col, startPos.line, direction)
				if !ok {
					continue OUTER
				}
				if !strings.Contains(toFind, string(wm.lines[posLine][posCol])) {
					continue OUTER
				}
				toFind = strings.Trim(toFind, string(wm.lines[posLine][posCol]))
			}
			if toFind != "" {
				continue OUTER
			}
		}
		findings++
	}

	return findings
}

func (wm *wordMap) move(col, line int, dir direction) (posCol, posLine int, ok bool) {
	offsets := moveMap[dir]
	posCol = col + offsets[0]
	posLine = line + offsets[1]
	if !(posCol < 0 || posLine < 0 || posCol >= len(wm.lines[0]) || posLine >= len(wm.lines)) {
		ok = true
	}
	return
}

func newWordMap(r io.Reader) *wordMap {
	line := 0
	var wm wordMap
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		currentline := scanner.Text()
		wm.lines = append(wm.lines, []byte(currentline))
		for col, ch := range currentline {
			if ch == 'X' {
				wm.startPosX = append(wm.startPosX, pos{line: line, col: col})
			}
			if ch == 'A' {
				wm.startPosA = append(wm.startPosA, pos{line: line, col: col})
			}
		}
		line++
	}
	return &wm

}

func main() {
	r, _ := os.Open("input.txt")
	wm := newWordMap(r)
	fmt.Printf("Part I: %d\n", wm.Search())
	fmt.Printf("Part II: %d\n", wm.SearchXMAS())
}
