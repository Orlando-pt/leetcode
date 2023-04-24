package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	head := ListNode{Val: 1}
	n1 := ListNode{Val: 2}
	head.Next = &n1

	n2 := ListNode{Val: 6}
	n1.Next = &n2

	n3 := ListNode{Val: 3}
	n2.Next = &n3

	n4 := ListNode{Val: 4}
	n3.Next = &n4

	n5 := ListNode{Val: 5}
	n4.Next = &n5

	n6 := ListNode{Val: 6}
	n5.Next = &n6

	// printLinkedList(&head)

	newList := removeElements(&head, 6)

	printLinkedList(newList)
}

func removeElements(head *ListNode, val int) *ListNode {

	// previous := head
	// for node := head; node != nil; node = node.Next {
	// 	if head.Val == val && head == node {
	// 		head = head.Next
	// 	}
	// 	if val == node.Val {
	// 		previous.Next = node.Next
	// 	}
	// 	previous = node
	// }
	//
	// return head

	// find head
	for head != nil && head.Val == val {
		head = head.Next
	}
	newHead := head

	if head == nil || head.Next == nil {
		return newHead
	}

	head = head.Next
	previousValid := newHead
	for head != nil {
		if head.Val != val {
			previousValid.Next = head
			if head.Next != nil {
				previousValid = previousValid.Next
			}
		}
		head = head.Next
	}

	if previousValid.Next != nil && previousValid.Next.Val == val {
		previousValid.Next = nil
	}

	return newHead

	// solution code
	// if head == nil {
	// 	return nil
	// }
	//
	// head.Next = removeElements(head.Next, val)
	// if head.Val == val {
	// 	return head.Next
	// }
	//
	// return head
}

func printLinkedList(head *ListNode) {
	temp := head
	for temp != nil {
		fmt.Println(temp.Val)
		temp = temp.Next
	}
}
