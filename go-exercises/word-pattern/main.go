package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(wordPattern("abba", "dog cat cat dog"))
}

func wordPattern(pattern string, s string) bool {
	matches := make(map[string]byte)
	patternIdx := 0
	for i, word := range strings.Split(s, " ") {
		if val, ok := matches[word]; ok {
			if val != pattern[i] {
				return false
			}
		} else {
			matches[word] = pattern[patternIdx]
		}
		patternIdx++
	}
	return true
}
