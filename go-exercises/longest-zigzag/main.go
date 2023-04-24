package main

import (
	"fmt"
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

	n1 := TreeNode{Val: 1}
	n2 := TreeNode{Val: 1}
	n3 := TreeNode{Val: 1}
	n4 := TreeNode{Val: 1}
	n5 := TreeNode{Val: 1}
	n6 := TreeNode{Val: 1}
    n7 := TreeNode{Val: 1}

	root.Right = &n1
	n1.Right = &n2
	n1.Left = &n3
	n2.Right = &n4
	n2.Left = &n5
	n5.Right = &n6
    n6.Right = &n7

	// printTree(&root, 0)
	fmt.Println(longestZigZag(&root))
}

func longestZigZag(root *TreeNode) int {
    // node2visit := make([]*TreeNode, 1)
    // biggestPath := 0
    //
    // node2visit[0] = root
    //
    // for len(node2visit) != 0 {
    //     curr := node2visit[0]
    //
    //     findLeft := findLongestPath(curr.Left, 1)
    //     findRight := findLongestPath(curr.Right, 0)
    //
    //     if findLeft > biggestPath {
    //         biggestPath = findLeft
    //     }
    //
    //     if findRight > biggestPath {
    //         biggestPath = findRight
    //     }
    //
    //     node2visit = node2visit[1:]
    //
    //     if curr.Left != nil {
    //         node2visit = append(node2visit, curr.Left)
    //     }
    //     if curr.Right != nil {
    //         node2visit = append(node2visit, curr.Right)
    //     }
    // }
    // return biggestPath
    _, _, res := dfs(root)
    return res
}

func dfs(root *TreeNode) (int, int, int) {

    if root == nil {
        return -1, -1, -1
    }

    _, leftR, leftRes := dfs(root.Left)
    rightL, _, rightRes := dfs(root.Right)

    res := getMax(getMax(leftR, rightL) + 1, getMax(leftRes, rightRes))

    return leftR + 1, rightL + 1, res
}

func getMax(n1 int, n2 int) int {
    if n1 > n2 {
        return n1
    }
    return n2
}

func findLongestPath(root *TreeNode, dir int) int {
    if root == nil {
        return 0
    }

    if dir == 0 {
        return 1 + findLongestPath(root.Left, 1)
    }

    return 1 + findLongestPath(root.Right, 0)
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
