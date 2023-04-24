package main

import "fmt"

const COUNT int = 10

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := TreeNode{Val: -2}

	n1 := TreeNode{Val: 2}
	n2 := TreeNode{Val: -3}
	n3 := TreeNode{Val: 4}
	n4 := TreeNode{Val: 5}

	// root.Left = &n1
	root.Right = &n2
	n1.Left = &n3
	n1.Right = &n4

	// printTree(&root, 8)
	fmt.Println(hasPathSum(&root, -5))
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil && root.Val == targetSum {
		return true
	}

	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)

	// return helperFunc(root, targetSum, 0)
}

func helperFunc(root *TreeNode, targetSum int, sum int) bool {

	if root == nil {
		return false
	}

	sum += root.Val

	if root.Left == nil && root.Right == nil && targetSum == sum {
		return true
	}

	return helperFunc(root.Left, targetSum, sum) || helperFunc(root.Right, targetSum, sum)
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
