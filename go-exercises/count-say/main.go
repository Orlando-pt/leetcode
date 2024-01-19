package main

import (
	"fmt"
	"strconv"
)

func main() {

    fmt.Println(countAndSay(1))
    fmt.Println(countAndSay(2))
}

func countAndSay(n int) string {
    number := "1"

    for i := 1; i < n; i++ {
        number = helper(number)
    }

    return number
}

func helper(curr string) string {
    result := ""
    i := 0
    for i < len(curr) {
        count := 1

        for i + 1 < len(curr) && curr[i] == curr[i+1] {
            i += 1
            count += 1
        }

        result += strconv.Itoa(count) + string(curr[i])
        i += 1
    }

    return result
}
