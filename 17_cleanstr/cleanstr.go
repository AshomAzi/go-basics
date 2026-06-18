package main

import (
	"fmt"
	"os"
	"strings"
)

func CleanStr(str string) {
	val := strings.Fields(str)
	for _, v := range val {
		fmt.Print(v+" ")
	}
	fmt.Println()
}

func main() {

	
	newStr := os.Args
	if len(newStr) != 2 {
		fmt.Println()
	} else {
		CleanStr(newStr[1])
	}

}

