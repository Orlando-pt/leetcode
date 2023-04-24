package main

import "fmt"

func main() {

	fmt.Println(singleNonDuplicate([]int{1, 1, 2, 3, 3, 4, 4, 8, 8}))
	fmt.Println(singleNonDuplicate([]int{3, 3, 7, 7, 10, 11, 11}))
}

func singleNonDuplicate(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	low := 0
	high := len(nums) - 1

	for low < high {
		middle := (low + high) / 2

		if (middle % 2) == 1 {
			middle--
		}

		if nums[middle] != nums[middle+1] {
			high = middle
		} else {
			low = middle + 2
		}
	}
	return nums[low]
}
