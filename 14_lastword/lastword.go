package main

import (
	"fmt"
	"strings"
)

func LastWord(str string) string {
	newStr := strings.Fields(str)
	finalStr := ""
	if len(newStr) > 0 {
		finalStr += newStr[len(newStr)-1]+"\n"
	} else {
		return ""
	}
	return finalStr
}

func main() {
	fmt.Print(LastWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Print(LastWord(" lorem,ipsum "))
	fmt.Print(LastWord(" "))
}