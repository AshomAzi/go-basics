package main

import (
	"fmt"
	"os"
	"strings"
)

func ExpandStr(str string) {
	val := strings.Join(strings.Fields(str), "   ")
	fmt.Println(val)
}

func main() {

	newStr := os.Args

	if len(newStr) != 2 {
		fmt.Println()
	} else {
		ExpandStr(newStr[1])
	}

}
