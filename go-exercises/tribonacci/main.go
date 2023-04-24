package main

import "fmt"

func main() {
	fmt.Println(tribonacci(4))
	fmt.Println(tribonacci(25))
	fmt.Println(tribonacci(36))
}

func tribonacci(n int) int {

	/**
	if n == 0 {
		return 0
	}

	if n < 3 {
		return 1
	}

	return tribonacci(n-1) + tribonacci(n-2) + tribonacci(n-3)
	*/

	if n == 0 {
		return 0
	}

	if n <= 2 {
		return 1
	}

	calculatedPositions := make([]int, n+1)
	calculatedPositions[0] = 0
	calculatedPositions[1] = 1
	calculatedPositions[2] = 1

	for i := 3; i <= n; i++ {
		calculatedPositions[i] = calculatedPositions[i-1] + calculatedPositions[i-2] + calculatedPositions[i-3]
	}

	return calculatedPositions[n]
}
