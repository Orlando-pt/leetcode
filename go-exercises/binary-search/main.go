package main

import (
	"fmt"
)

func main() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9))
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 2))
}

func search(nums []int, target int) int {

	left, right := 0, len(nums)-1

	for left <= right {
		middle := left + (right-left)/2

		if nums[middle] == target {
            return middle
		} else if nums[middle] < target {
            left = middle + 1
		} else {
            right = middle - 1
		}
	}

	return -1
}
