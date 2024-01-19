package main

import "fmt"

func main() {
	fmt.Println(strStr("hello", "ll"))
	fmt.Println(strStr("aaaaa", "bba"))
	fmt.Println(strStr("aaaaa", "a"))
	fmt.Println(strStr("aaaaa", "aa"))
    fmt.Println(strStr("sadbutsad", "sad"))
    fmt.Println(strStr("leetcode", "leeto"))
    fmt.Println(strStr("mississippi", "issip"))
}

func strStr(haystack string, needle string) int {
    if len(needle) > len(haystack) {
        return - 1
    }

	for i := 0; i < len(haystack); i++ {
        if haystack[i] == needle[0] {
            for n := 0; n < len(needle); n++ {
                if i+n > len(haystack) || haystack[i+n] != needle[n] {
                    break
                }

                if n == len(needle) - 1 {
                    return i
                }
            }
        }
	}
	return -1
}
