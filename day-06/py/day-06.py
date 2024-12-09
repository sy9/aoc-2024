# Day       Time   Rank  Score       Time   Rank  Score
#   6   00:04:42    277      0   00:41:35   2622      0

import time

class T(tuple):
    def __add__(self, other):
        return T(x + y for x, y in zip(self, other))

    def rot(self):
        x, y = self
        return T((y, -x))


NORTH = T((-1, 0))


def sim(grid, pos, dir):
    visited = set()
    visitedp = set()
    while pos in grid:
        if (pos, dir) in visited:
            return True, visited
        visited.add((pos, dir))
        visitedp.add(pos)
        if len(visited) > 3200:
            printg(grid, visitedp)
        while grid.get(pos + dir, ".") in ("#", "O"):
            dir = dir.rot()
        pos += dir
    return False, visited

def printg(grid, visited):
    # Get the number of rows and columns from the grid
    max_row = max(i for i, j in grid.keys()) + 1
    max_col = max(j for i, j in grid.keys()) + 1

    # Reconstruct the grid and print it
    for i in range(max_row):
      row = ''
      for j in range(max_col):
        c = grid.get((i, j))
        if c == '.':
            c = ' '
        if (i,j) in visited:
            c = '.'
        row += c
      print(row)
    time.sleep(1)

       #row = ''.join(' ' if grid.get((i, j), ' ') == '.' else grid.get((i, j), ' ') for j in range(max_col))  # default to ' ' if missing
       #print(row)

def p1(f):
    grid = {
        T((i, j)): c
        for i, r in enumerate(f.read().splitlines())
        for j, c in enumerate(r)
    }
   
   

    pos = next(p for p, c in grid.items() if c == "^")
    _, visited = sim(grid, pos, NORTH)
    visited = {p for p, d in visited}
    return len(visited)


def p2(f):
    grid = {
        T((i, j)): c
        for i, r in enumerate(f.read().splitlines())
        for j, c in enumerate(r)
    }
    pos = next(p for p, c in grid.items() if c == "^")
    #_, visited = sim(grid, pos, NORTH)
    #visited = {p for p, d in visited}
    visited = set()
    visited.add((70,39))
    # print(pos)

    ans = 0

    for vpos in visited:
        if grid[vpos] == "#":
            continue
        grid[vpos] = "O"
        works, _ = sim(grid, pos, NORTH)
        if works:
            ans += works
            #print(vpos)
        grid[vpos] = "."

    return ans

#print(p1(open("../input.txt")))
print(p2(open("../input.txt")))