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
	fmt.Println(preorderTraversal(&root))
}

func preorderTraversal(root *TreeNode) []int {

	// res := make([]int, 0)
	// return helperFunction(root, res)

	return stackImplementation(root)
}

func helperFunction(root *TreeNode, res []int) []int {
	if root == nil {
		return res
	}

	res = helperFunction(root.Left, res)
	res = helperFunction(root.Right, res)
	res = append(res, root.Val)
	return res
}

func stackImplementation(root *TreeNode) []int {
	res := make([]int, 0)

	var stack []*TreeNode

	stack = append(stack, root)

	for len(stack) > 0 {
		node := stack[len(stack)-1]

		if node == nil {
			stack = stack[:len(stack)-1]
			continue
		}

		res = append(res, node.Val)

		stack = stack[:len(stack)-1]
		stack = append(stack, node.Left)
		stack = append(stack, node.Right)
	}

	return reverseArray(&res)
}

func reverseArray(array *[]int) []int {
	for i, j := 0, len(*array)-1; i < j; i, j = i+1, j-1 {
		(*array)[i], (*array)[j] = (*array)[j], (*array)[i]
	}

	return *array
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
