package main

import "fmt"

func main() {
	fmt.Println(canConstruct("aa", "aabz"))
	fmt.Println(canConstruct("aaaz", "aabz"))
}

func canConstruct(ransomNote string, magazine string) bool {
	letters := make([]int, 26)
	for _, c := range magazine {
		pos := c - 'a'
		letters[pos]++
	}

	for _, c := range ransomNote {
		pos := c - 'a'
		letters[pos]--

		if letters[pos] < 0 {
			return false
		}
	}

	return true
}
