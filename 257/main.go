package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("!!")

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	var ans []string
	var path []int
	dfs(root, &path, &ans)
	return ans
}

func dfs(node *TreeNode, path *[]int, ans *[]string) {
	if node == nil {
		return
	}
	*path = append(*path, node.Val)

	if node.Left == nil && node.Right == nil {
		strPath := make([]string, len(*path))

		for i, val := range *path {
			strPath[i] = strconv.Itoa(val)
		}

		out := strings.Join(strPath, "->")
		*ans = append(*ans, out)
	}
	dfs(node.Left, path, ans)
	dfs(node.Right, path, ans)

	*path = (*path)[:len(*path)-1]
}
