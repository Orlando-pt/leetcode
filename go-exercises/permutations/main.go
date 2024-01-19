package main

import "fmt"

func main() {

	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	res := make([][]int, 0)

	backtrack(&res, make([]int, 0), nums)
	return res
}

func backtrack(list *[][]int, tempList []int, nums []int) {
	if len(tempList) == len(nums) {
		dest := make([]int, len(tempList))
		copy(dest, tempList)

		*list = append(*list, dest)
	} else {
		for i := 0; i < len(nums); i++ {
			if contains(tempList, nums[i]) {
				continue
			}

			tempList = append(tempList, nums[i])
			backtrack(list, tempList, nums)
			tempList = tempList[:len(tempList)-1]
		}
	}
}

func contains(list []int, elem int) bool {
	for _, e := range list {
		if e == elem {
			return true
		}
	}

	return false
}
