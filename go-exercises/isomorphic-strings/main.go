package main

import "fmt"

func main() {
	fmt.Println(isIsomorphic("egg", "add"))
	fmt.Println(isIsomorphic("foo", "bar"))
	fmt.Println(isIsomorphic("paper", "title"))
}

func isIsomorphic(s string, t string) bool {
	sMap := make(map[byte]byte)
	tMap := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		val, ok := sMap[s[i]]
		val2, ok2 := tMap[t[i]]

		if ok && val != t[i] {
			return false
		}
		if ok2 && val2 != s[i] {
			return false
		} else if !ok && !ok2 {
			sMap[s[i]] = t[i]
			tMap[t[i]] = s[i]
		}
	}
	return true
}
