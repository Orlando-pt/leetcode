package main

import "fmt"

func main() {
	//fmt.Println(intToRoman(3))
	//fmt.Println(intToRoman(58))
	//fmt.Println(intToRoman(1994))

	fmt.Println(intToRomanBest(3))
	fmt.Println(intToRomanBest(58))
	fmt.Println(intToRomanBest(1994))
}

type Roman struct {
	S string
	V int
}

func intToRomanBest(num int) string {
	ans := ""
	len := 12

	r := []Roman{
		{"I", 1},
		{"IV", 4},
		{"V", 5},
		{"IX", 9},
		{"X", 10},
		{"XL", 40},
		{"L", 50},
		{"XC", 90},
		{"C", 100},
		{"CD", 400},
		{"D", 500},
		{"CM", 900},
		{"M", 1000},
	}

	for i := len; i >= 0; i-- {

		if times := num / r[i].V; times != 0 {

			if times == 1 {
				ans += r[i].S
			} else {
				for j := 0; j < times; j++ {
					ans += r[i].S
				}
			}

			num = num % r[i].V
		}
	}
	return ans
}

func intToRoman(num int) string {
	mappings := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}

	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	res := ""

	for num > 0 {
		for _, n := range nums {
			subtraction := num - n
			if subtraction >= 0 {
				res += mappings[n]
				num = subtraction
				break
			}
		}
	}
	return res
}
