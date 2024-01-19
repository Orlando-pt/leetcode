package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

func groupAnagrams(strs []string) [][]string {
	strsMap := make(map[string][]string)

	for _, v := range strs {
		sortedStr := sortString(v)

		if _, ok := strsMap[sortedStr]; ok {
			strsMap[sortedStr] = append(strsMap[sortedStr], v)
		} else {
			strsMap[sortedStr] = []string{v}
		}
	}

    res := make([][]string, len(strsMap))

    i := 0
    for _, anagrams := range strsMap {
        res[i] = anagrams
        i++
    }

	return res
}

func sortString(value string) string {
	strRune := []rune(value)

	sort.Slice(strRune, func(i, j int) bool {
		return strRune[i] < strRune[j]
	})

	return string(strRune)
}

