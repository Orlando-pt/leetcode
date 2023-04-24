package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "    fly me to the moon    "

	words := strings.Fields(s)
	fmt.Println(len(words[len(words)-1]))
}
