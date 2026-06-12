package main

import (
	"fmt"
	"strings"
)

func FirstWord(str string) string {
	new := strings.Split(str, " ")
	return new[0]+"\n"
}

func main() {
    fmt.Print(FirstWord("hello there"))
    fmt.Print(FirstWord(""))
    fmt.Print(FirstWord("hello   .........  bye"))
}
