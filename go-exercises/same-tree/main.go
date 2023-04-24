package main

import "fmt"

const COUNT int = 10

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := TreeNode{Val: 1}

	n1 := TreeNode{Val: 2}
	n2 := TreeNode{Val: 3}
	n3 := TreeNode{Val: 4}
	n4 := TreeNode{Val: 5}

	root.Left = &n1
	root.Right = &n2
	n1.Left = &n3
	n1.Right = &n4

	//printTree(&root, 0)
	fmt.Println(isSameTree(&root, &n1))
}

func isSameTree(p *TreeNode, q *TreeNode) bool {

	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	}

	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func printTree(root *TreeNode, space int) {
	if root == nil {
		return
	}

	space += COUNT

	printTree(root.Right, space)

	fmt.Print("\n")
	for i := COUNT; i < space; i++ {
		fmt.Print(" ")
	}
	fmt.Println(root.Val)

	printTree(root.Left, space)
}
