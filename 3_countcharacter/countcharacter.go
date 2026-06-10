package main

import "fmt"

func CountCharacter(input string, char rune) int {
	count := 0

	if len(input) == 0 {
		return 0
	}

	for _, v := range input {
		if v == char {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(CountCharacter("HellHo", 'H'))
}
