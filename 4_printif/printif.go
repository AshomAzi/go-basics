package main

import "fmt"

func PrintIf(input string) string {

	if input == "" {
		return "G\n"
	}

	if len(input) < 3 {
		return "Invalid Input \n"
	}
	return "G\n"
}

func main() {
	fmt.Println(PrintIf("Hellp"))
}