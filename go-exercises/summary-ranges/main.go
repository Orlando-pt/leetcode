package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
	fmt.Println(summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
}

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	res := make([]string, 0)

	first := nums[0]
	prev := nums[0]
	for i := 1; i < len(nums); i++ {
		curr := nums[i]
		if curr != prev+1 {
			if first == prev {
				res = append(res, strconv.Itoa(first))
			} else {
				res = append(res, strconv.Itoa(first)+"->"+strconv.Itoa(prev))
			}
			first = curr
		}
		prev = curr
	}

	if first == prev {
		res = append(res, strconv.Itoa(first))
	} else {
		res = append(res, strconv.Itoa(first)+"->"+strconv.Itoa(prev))
	}
	return res
}
