package main

import "fmt"

func main() {
	a := "1010"
	b := "1011"
	res := ""
	reminder := false

	a_curr := len(a) - 1
	b_curr := len(b) - 1

	var a_value = ""
	var b_value = ""

	var i int
	if a_curr > b_curr {
		i = a_curr
	} else {
		i = b_curr
	}

	for i >= 0 {

		if a_curr > -1 {
			a_value = string(a[a_curr])
			a_curr--
		} else {
			a_value = "0"
		}

		if b_curr > -1 {
			b_value = string(b[b_curr])
			b_curr--
		} else {
			b_value = "0"
		}

		if a_value == "1" && b_value == "1" {
			if reminder {
				res = "1" + res
			} else {
				res = "0" + res
			}
			reminder = true
		} else if a_value == "0" && b_value == "0" {
			if reminder {
				res = "1" + res
			} else {
				res = "0" + res
			}
			reminder = false
		} else {
			if reminder {
				res = "0" + res
				reminder = true
			} else {
				res = "1" + res
				reminder = false
			}
		}

		i--
	}

	if reminder {
		res = "1" + res
	}

	fmt.Println(res)
}
