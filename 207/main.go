package main

import (
	"fmt"
	"slices"
)

func main() {
	tests := []struct {
		numCourses    int
		prerequisites [][]int
	}{
		{
			numCourses: 2,
			prerequisites: [][]int{
				{1, 0},
			},
		},
		{
			numCourses: 2,
			prerequisites: [][]int{
				{1, 0},
				{0, 1},
			},
		},
		{
			numCourses: 4,
			prerequisites: [][]int{
				{1, 0},
				{2, 0},
				{3, 1},
				{3, 2},
			},
		},
	}

	for _, t := range tests {
		fmt.Println(canFinish(t.numCourses, t.prerequisites))
	}
}

func canFinish(numCourses int, prerequisites [][]int) bool {

	visited := make([]bool, numCourses)
	inPath := make([]bool, numCourses)
	graph := initializeGraph(numCourses, prerequisites)
	var dfs func(v int) bool
	dfs = func(v int) bool {

		if inPath[v] {
			return true
		}
		if visited[v] {
			return false
		}
		visited[v], inPath[v] = true, true
		if slices.ContainsFunc(graph[v], dfs) {
			return true
		}
		inPath[v] = false

		return false
	}

	for v := range numCourses {
		if dfs(v) {
			return false
		}
	}
	return true
}

func initializeGraph(numCourses int, prerequisites [][]int) [][]int {
	graph := make([][]int, numCourses)
	for _, p := range prerequisites {
		graph[p[1]] = append(graph[p[1]], p[0])
	}
	return graph
}
