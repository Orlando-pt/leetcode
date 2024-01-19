package main

import "fmt"

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
	fmt.Println(search([]int{1}, 0))
}

func search(nums []int, target int) int {
	lo, hi := 0, len(nums)-1

	for lo < hi {
		mid := (lo + hi) / 2

		if nums[mid] > nums[hi] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	rot := lo
	lo, hi = 0, len(nums)-1

	for lo <= hi {
		mid := (lo + hi) / 2

		realmid := (mid + rot) % len(nums)

		if nums[realmid] == target {
			return realmid
		}

		if nums[realmid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return -1
}
