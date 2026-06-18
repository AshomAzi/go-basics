package main

import (
	"fmt"
	"os"
	"strings"
)

func Replace(str, val1, val2 string) string {
	if !strings.Contains(str, val1) {
		return str
	} else {
		return strings.ReplaceAll(str, val1, val2)
	}
}

func main() {
	newStr := os.Args
	if len(newStr) == 4 {
		fmt.Println(Replace(newStr[1], newStr[2], newStr[3]))
	} else {
		fmt.Println()
	}
}
