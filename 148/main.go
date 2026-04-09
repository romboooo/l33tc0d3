package main

import "fmt"

func main() {
	head := &ListNode{4, &ListNode{2, &ListNode{1, &ListNode{3, nil}}}}

	res := sortList(head)
	printList(res)
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}
