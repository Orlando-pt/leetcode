package main

import "fmt"

func main() {
	fmt.Println(convertToTitle(1))
	fmt.Println(convertToTitle(28))
	fmt.Println(convertToTitle(701))
}

func convertToTitle(n int) string {
	var result string
	for n > 0 {
		n--
		result = string(n%26+'A') + result
		n /= 26
	}
	return result
}
