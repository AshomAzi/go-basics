package main

import "fmt"

func CountAlpha(input string) int {
	count := 0
	for _, v := range input {
		if (v >= 'A' && v <= 'Z') || v >= 'a' && v <= 'z' {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(CountAlpha("1 2 3 4 5 A  b A B C"))
}