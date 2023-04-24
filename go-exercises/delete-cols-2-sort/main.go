package main

import "fmt"

func main() {

	fmt.Println(minDeletionSize([]string{"a", "b"}))
}

func minDeletionSize(strs []string) int {
	count := 0
	for i := 0; i < len(strs[0]); i++ {
		for j := 0; j < len(strs)-1; j++ {
			if strs[j][i] > strs[j+1][i] {
				count++
				break
			}
		}
	}
	return count
}
