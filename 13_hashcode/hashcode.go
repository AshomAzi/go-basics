package main

import "fmt"

func HashCode(str string) string {
	newStr := ""
	for _, i := range str {
		if i < 32 || i > 126 {
			i = i+33
		}
		newStr+=string((i+int32(len(str)))%127)
	}
	return newStr
}

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello😊 World"))
}
