package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{1, 1}))
}

func maxArea(height []int) int {
	max := 0
	head := 0
	tail := len(height) - 1

	for head < tail {

		min := getMin(height[head], height[tail])

		area := min * (tail - head)

		if area > max {
			max = area
		}

        if height[head] > height[tail] {
            tail--
        } else {
            head++
        }
	}

	return max
}

func getMin(n1 int, n2 int) int {
	if n1 < n2 {
		return n1
	}

	return n2
}
