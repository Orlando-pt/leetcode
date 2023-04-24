package main

import (
	"fmt"
	"math"
)

func main() {
    fmt.Println(myAtoi("4193 with words"))
    fmt.Println(myAtoi("words and 987"))
    fmt.Println(myAtoi("00000-42a1234"))
}

func myAtoi(s string) int {
    sign := 1
    base := 0
    i := 0

    for i < len(s) && s[i] == ' ' {
        i++
    }

    if i < len(s) && (s[i] == '-' || s[i] == '+') {
        if i < len(s) && s[i] == '-' {
            sign = -1
        }
        i++
    }


    for i < len(s) && s[i] >= '0' && s[i] <= '9' {
        if base > math.MaxInt32 / 10 || (base == math.MaxInt32 / 10 && s[i] - '0' > 7) {
            if sign == 1 {
                return math.MaxInt32
            }
            return math.MinInt32
        }

        base = 10 * base + int(s[i] - '0')
        i++
    }

    return base * sign
}

func myAtoiNotWorking(s string) int {
    var zero rune = '0'
    var nine rune = '9'
    var space rune = ' '
    var dash rune = '-'
    var plus rune = '+'

    positive := true
    signIsAllowed := true
    spaceValidation := false
    numbers := 0

    var ans int32 = 0

    for pos, char := range s {
        if char == space {
            spaceValidation = true
            continue
        }

        if signIsAllowed && (char == dash || char == plus) {
            if char == dash {
                positive = false
            }

            signIsAllowed = false
            continue
        }

        if char >= zero && char <= nine {
            parsedInt := parseIntChar(char)

            newAns := ans * 10 + parsedInt

            if (newAns - parsedInt) / 10 != ans {
                if newAns < 0 {
                    return math.MaxInt32
                }

                return math.MinInt32
            }
            ans = newAns
            spaceValidation = false
            numbers++
        } else {
            if spaceValidation || numbers == pos {
                return returnInt32(positive, ans)
            }
            return 0
        }
    }

    return returnInt32(positive, ans)
}

func returnInt32(positive bool, ans int32) int{
    if !positive {
        return -int(ans)
    }
    return int(ans)
}

func parseIntChar(c rune) int32 {
    var zero rune = '0'
    ans := c - zero

    return int32(ans)
}
