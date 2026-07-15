package main

func main() {

}

func numIslands(grid [][]byte) int {

	lenY := len(grid)
	lenX := len(grid[0])
	count := 0
	var dfs func(x, y int)
	dfs = func(x, y int) {

		if x < 0 || x > lenX-1 || y < 0 || y > lenY-1 {
			return
		}

		if grid[y][x] == '0' {
			return
		}
		grid[y][x] = '0'

		dfs(x+1, y)
		dfs(x-1, y)
		dfs(x, y+1)
		dfs(x, y-1)
	}

	for x := range lenX {
		for y := range lenY {
			if grid[y][x] == '1' {
				count++
				dfs(x, y)
			}
		}
	}

	return count
}
