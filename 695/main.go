package main

import "fmt"

func main() { fmt.Println("!") }

func maxAreaOfIsland(grid [][]int) int {
	lenY := len(grid)
	lenX := len(grid[0])

	maxArea := 0

	var dfs func(x, y int) int
	dfs = func(x, y int) int {
		area := 0
		if x >= lenX || y >= lenY || x < 0 || y < 0 {
			return area
		}

		if grid[y][x] == 0 {
			return area
		}
		if grid[y][x] == 1 {
			grid[y][x] = 0
			area++
		}

		area += dfs(x+1, y)
		area += dfs(x-1, y)
		area += dfs(x, y-1)
		area += dfs(x, y+1)

		return area
	}

	for x := range lenX {
		for y := range lenY {

			if grid[y][x] == 1 {
				area := dfs(x, y)

				if maxArea < area {
					maxArea = area
				}

			}

		}
	}

	return maxArea
}
