package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := ListNode{Val: 1}

	n1 := ListNode{Val: 1}
	head.Next = &n1

	n2 := ListNode{Val: 2}
	n1.Next = &n2

	n3 := ListNode{Val: 3}
	n2.Next = &n3

	n4 := ListNode{Val: 3}
	n3.Next = &n4

	printList(deleteDuplicates(&head))
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	currentHead := head

	temp := head.Next
	for temp != nil {
		if temp.Val == currentHead.Val {
			temp = temp.Next
		} else {
			currentHead.Next = temp
			currentHead = temp
			temp = temp.Next
		}
	}

	if currentHead.Next != nil && currentHead.Val == currentHead.Next.Val {
		currentHead.Next = nil
	}

	return head
}

func printList(head *ListNode) {
	temp := head
	for temp != nil {
		fmt.Println(temp.Val)
		temp = temp.Next
	}
}
