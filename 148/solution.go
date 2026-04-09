package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func merge(firstHead *ListNode, secondHead *ListNode) *ListNode {

	tempNode := &ListNode{}
	curr := tempNode

	for firstHead != nil && secondHead != nil {

		if firstHead.Val > secondHead.Val {
			curr.Next = secondHead
			secondHead = secondHead.Next

		} else {
			curr.Next = firstHead
			firstHead = firstHead.Next

		}
		curr = curr.Next
	}
	if firstHead != nil {
		curr.Next = firstHead
	}

	if secondHead != nil {
		curr.Next = secondHead
	}
	return tempNode.Next
}

func sortList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	slowPointer := head
	fastPointer := head.Next

	for fastPointer != nil && fastPointer.Next != nil {
		slowPointer = slowPointer.Next
		fastPointer = fastPointer.Next.Next
	}

	rightHead := slowPointer.Next
	slowPointer.Next = nil

	sortedLeft := sortList(head)
	sortedRight := sortList(rightHead)

	return merge(sortedLeft, sortedRight)
}
