package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
	fmt.Println(isPalindrome(" "))
	fmt.Println(isPalindrome("ab_a"))
}

func isPalindrome(s string) bool {

	left := 0
	right := len(s) - 1

	for left != right {
		if !isAlphanumeric(s[left]) {
			left++
			continue
		}
		if !isAlphanumeric(s[right]) {
			right--
			continue
		}

		if toUpper(s[left]) != toUpper(s[right]) {
			return false
		}
		left++
		right--
	}

	return true
}

func toUpper(c byte) byte {
	if c >= 'a' && c <= 'z' {
		return c - 0x20
	}

	return c
}

func isAlphanumeric(c byte) bool {
	return (c >= '0' && c <= '9') || (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z')
}
