package main

import "fmt"

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
	fmt.Println(containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))

	fmt.Println(mostVotedAnswer([]int{1, 2, 3, 1}, 3))
	fmt.Println(mostVotedAnswer([]int{1, 0, 1, 1}, 1))
	fmt.Println(mostVotedAnswer([]int{1, 2, 3, 1, 2, 3}, 2))
}

func containsNearbyDuplicate(nums []int, k int) bool {
	visitedIndexes := make(map[int]int)
	for i, num := range nums {
		position, ok := visitedIndexes[num]

		if ok {
			if i-position <= k {
				return true
			}
		}
		visitedIndexes[num] = i
	}
	return false
}

func quickest(nums []int, k int) bool {
	seen := make(map[int]int, len(nums))

	for i, num := range nums {
		if pos, ok := seen[num]; ok && i-pos <= k {
			return true
		}

		seen[num] = i
	}

	return false
}

type void struct{}

var nothing void

func mostVotedAnswer(nums []int, k int) bool {
	set := make(map[int]void)

	for i := 0; i < len(nums); i++ {
		if i > k {
			delete(set, nums[i-k-1])
		}
		if _, contains := set[nums[i]]; contains {
			return true
		} else {
			set[nums[i]] = nothing
		}
	}
	return false
}
