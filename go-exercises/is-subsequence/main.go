package main

import "fmt"

func main() {
	fmt.Println(isSubsequence("abc", "ahbgdc"))
	fmt.Println(isSubsequence("axc", "ahbgdc"))
	fmt.Println(isSubsequence("b", "abc"))
}

func isSubsequence(s string, t string) bool {
	if s == "" {
		return true
	}

	if t == "" {
		return false
	}

	for i := range t {
		if t[i] == s[0] {
			ii := i + 1
			subToCompare := 1

			if subToCompare == len(s) {
				return true
			}

			for ii < len(t) {

				if t[ii] != s[subToCompare] {
					ii++
				} else {
					ii++
					subToCompare++
				}

				if subToCompare == len(s) {
					return true
				}
			}
		}
	}

	return false
}
