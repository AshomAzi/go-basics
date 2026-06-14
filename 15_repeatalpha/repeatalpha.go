package main

import (
	"fmt"
	"strings"
)

func RepeatAlpha(str string) string {

	newStr := ""
	alpha := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for _, v := range str {
		val := string(v)
		if strings.Contains(alpha, strings.ToUpper(val)) {
			for j, h := range alpha {
				if string(h) == strings.ToUpper(val) {
					newStr += strings.Repeat(string(v), j+1)
				}
			}
		}
	}

	return newStr
}

func main() {
	fmt.Println(RepeatAlpha("abc"))
	fmt.Println(RepeatAlpha("Choumi."))
	fmt.Println(RepeatAlpha(""))
	fmt.Println(RepeatAlpha("abacadaba 01!"))
}
