package main

import "fmt"

func main() {
	// head := &ListNode{4, &ListNode{2, &ListNode{1, &ListNode{3, nil}}}}

	// res := sortList(head)
	// printList(res)

	arr := []int{5, 2, 9, 1}

	fmt.Println(arr)
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}
