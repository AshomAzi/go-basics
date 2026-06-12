package main

import "fmt"

func FishAndChips(num int) string {
	if num < 0 {
		return "Error: Number is negative."
	}

	if num%2 != 0 && num%3 != 0 {
		return "Error: Non divisible"
	}
	if num%2 == 0 && num%3 == 0 {
		return "fish and chips"
	}

	if num%2 == 0 {
		return "fish"
	} else {
		return "chips"
	}
}

func main() {
	fmt.Println(FishAndChips(4))
	fmt.Println(FishAndChips(9))
	fmt.Println(FishAndChips(6))
}
