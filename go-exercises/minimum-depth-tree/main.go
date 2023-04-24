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

	printTree(&root, 0)
	fmt.Println(minDepth(&root))
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := minDepth(root.Left)
	right := minDepth(root.Right)

	if left == 0 || right == 0 {
		return left + right + 1
	} else {
		if left < right {
			return left + 1
		} else {
			return right + 1
		}
	}
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
