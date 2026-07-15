package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 5,
							},
						},
					},
				},
			},
		},
	}

	head = deleteDuplicates(head)

	printList(head)
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print(" -> ")
		}
		head = head.Next
	}
	fmt.Println()
}

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy
	curr := head

	for curr != nil {
		if curr.Next != nil && curr.Val == curr.Next.Val {
			dupVal := curr.Val
			for curr != nil && curr.Val == dupVal {
				curr = curr.Next
			}
			prev.Next = curr
		} else {
			prev = curr
			curr = curr.Next
		}
	}
	return dummy.Next
}
