package main

import "fmt"

func CheckNumber(input string) bool {

	for _, i := range input {
		if i >= '0' && i <= '9' {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(CheckNumber("Hello, World"))
}