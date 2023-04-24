package main

import "fmt"

func main() {
	fmt.Println(isPowerOfTwo(1))
	fmt.Println(isPowerOfTwo(16))
	fmt.Println(isPowerOfTwo(3))
}

func isPowerOfTwo(n int) bool {
	/**
	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}

	currVal := 2
	for currVal <= n {
		if currVal == n {
			return true
		}

		currVal *= 2
	}
	return false
	*/
	if n == 0 {
		return false
	}
	return (n & (n - 1)) == 0
}
