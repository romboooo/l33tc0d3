package main

func main() {
}

type Coordinates struct {
	x int
	y int
}

func orangesRotting(grid [][]int) int {

	dirs := [][]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}

	lenY := len(grid)
	lenX := len(grid[0])

	queue := []Coordinates{}

	freshOrangesCount := 0
	for x := range lenX {
		for y := range lenY {
			if grid[y][x] == 2 {
				queue = append(queue, Coordinates{x: x, y: y})
			}
			if grid[y][x] == 1 {
				freshOrangesCount++
			}
		}
	}
	minutes := 0
	for len(queue) > 0 {
		size := len(queue)
		for i := range size {
			currCoords := queue[i]

			for _, dir := range dirs {
				dx := dir[0]
				dy := dir[1]

				nx := currCoords.x + dx
				ny := currCoords.y + dy

				if nx < 0 || nx >= lenX || ny < 0 || ny >= lenY {
					continue
				}
				if grid[ny][nx] == 1 {
					grid[ny][nx] = 2
					freshOrangesCount--
					queue = append(queue, Coordinates{x: nx, y: ny})
				}
			}
		}
		queue = queue[size:]
		if len(queue) > 0 {
			minutes++
		}
	}

	if freshOrangesCount > 0 {
		return -1
	}

	return minutes
}
