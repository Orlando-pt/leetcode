package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	fmt.Println(searchRange([]int{}, 0))
}

func searchRange(nums []int, target int) []int {
	left := 0
	right := len(nums) - 1

	targetPosition := -1
	for left <= right {
		middle := int(math.Floor(float64((left + right) / 2)))

		if nums[middle] < target {
			left = middle + 1
		} else if nums[middle] > target {
			right = middle - 1
		} else {
			targetPosition = middle
            break
		}
	}

	if targetPosition == -1 {
		return []int{-1, -1}
	}

	// check left
    leftPosTarget := targetPosition
    for leftPosTarget >= 0 && nums[leftPosTarget] == target {
        leftPosTarget--
    }

    // check right
    for targetPosition < len(nums) && nums[targetPosition] == target {
        targetPosition++
    }

    return []int{leftPosTarget+1, targetPosition-1}
}
