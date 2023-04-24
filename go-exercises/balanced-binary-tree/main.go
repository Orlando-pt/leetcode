package main

import (
	"fmt"
	"math"
)

const COUNT int = 10

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type void struct{}

var member void

func main() {
	root := TreeNode{Val: 1}

	n1 := TreeNode{Val: 2}
	n2 := TreeNode{Val: 3}
	n3 := TreeNode{Val: 4}
	n4 := TreeNode{Val: 5}
	n5 := TreeNode{Val: 6}
	n6 := TreeNode{Val: 7}

	root.Left = &n1
	root.Right = &n2
	n1.Left = &n3
	n1.Right = &n4
	n3.Left = &n5
	n3.Right = &n6

	printTree(&root, 0)
	isBalanced(&root)
}

func isBalanced(root *TreeNode) bool {
	return checkBalanced(root) != -1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func checkBalanced(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftTree := checkBalanced(root.Left)
	if leftTree == -1 {
		return -1
	}

	rightTree := checkBalanced(root.Right)
	if rightTree == -1 {
		return -1
	}

	if abs(leftTree-rightTree) > 1 {
		return -1
	}

	if leftTree > rightTree {
		return leftTree + 1
	}

	return rightTree + 1
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
