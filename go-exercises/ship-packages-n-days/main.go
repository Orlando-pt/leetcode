package main

import "fmt"

func main() {
	fmt.Println(shipWithinDays([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5))
}

func shipWithinDays(weights []int, days int) int {
	// Find the max weight
	max := 0
	for _, w := range weights {
		if w > max {
			max = w
		}
	}

	// Find the sum of all weights
	sum := 0
	for _, w := range weights {
		sum += w
	}

	// Find the minimum capacity
	min := max
	for i := max; i <= sum; i++ {
		if canShip(weights, days, i) {
			min = i
			break
		}
	}

	return min
}

func canShip(weights []int, days int, capacity int) bool {
	// Iterate through the weights
	d := 1
	c := 0
	for _, w := range weights {
		if c+w > capacity {
			d++
			c = 0
		}
		c += w
	}

	return d <= days
}
