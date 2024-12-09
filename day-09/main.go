package main

import (
	"fmt"
	"os"
	"slices"
)

// block stores id, -1 == free
type block int

func parse(line string) []block {
	id := 0
	isFreeSpace := false
	disk := make([]block, 0)
	for _, input := range line {
		count := input - '0'
		blockId := -1
		if !isFreeSpace {
			blockId = id
			id++
		}
		for range count {
			disk = append(disk, block(blockId))
		}
		isFreeSpace = !isFreeSpace
	}
	return disk
}

func move(disk []block) []block {
	freeSpaceIdx := 0
	for i := len(disk) - 1; i >= 0; i-- {
		freeSpaceIdx = slices.Index(disk, -1)
		if freeSpaceIdx == -1 {
			break
		}
		disk[freeSpaceIdx], disk[i] = disk[i], disk[freeSpaceIdx]
		disk = slices.Delete(disk, i, i+1)
	}
	return disk
}

func moveP2(disk []block) []block {
	seen := make(map[int]struct{})
	fsm := newFreeSpaceMap(disk)
	size := 0
	fileID := -1
	for i := len(disk) - 1; i >= 0; i-- {
		if (disk[i] == -1 || int(disk[i]) != fileID) && size > 0 {
			if _, ok := seen[fileID]; ok {
				size = 0
			} else {
				seen[fileID] = struct{}{}
				targetIdx := fsm.queryLeftmost(size, i+1)
				if targetIdx == -1 {
					size = 0
				} else {
					for j := range size {
						disk[targetIdx+j], disk[i+j+1] = disk[i+j+1], disk[targetIdx+j]
					}
					size = 0
				}
			}
		}
		if disk[i] != -1 {
			fileID = int(disk[i])
			size++
		}
	}
	return disk
}

func checksum(disk []block) int {
	sum := 0
	for i := range disk {
		if disk[i] == -1 {
			continue // ignore free space blocks
		}
		sum += i * int(disk[i])
	}
	return sum
}

type freeSpaceMap struct {
	m map[int][]int // key = size, val = indices
	o [][2]int      // ordered list of size, index free spaces
}

func (f *freeSpaceMap) queryLeftmost(size, idx int) int {
	for i, fs := range f.o {
		if fs[0] < size {
			continue
		}
		if fs[1] > idx {
			break
		}
		fs[0] -= size
		fs[1] += size
		f.o[i] = fs
		return fs[1] - size
	}
	return -1
}

func newFreeSpaceMap(disk []block) *freeSpaceMap {
	fsm := &freeSpaceMap{
		m: make(map[int][]int),
	}
	size := 0
	for i := range disk {
		if disk[i] != -1 && size != 0 {
			fsm.o = append(fsm.o, [2]int{size, i - size})
			indices := fsm.m[size]
			indices = append(indices, i-size)
			fsm.m[size] = indices
			size = 0
			continue
		}
		if disk[i] == -1 {
			size++
		}
	}

	return fsm
}

func main() {
	disk, _ := os.ReadFile("input.txt")
	fmt.Printf("Part A: %d\n", checksum(move(parse(string(disk)))))
	fmt.Printf("Part B: %d\n", checksum(moveP2(parse(string(disk)))))
}
