package main

import "fmt"

func DigitLen(num1, num2 int) int {

	if num1 < 0 {
		num1 = -num1
	}

	if num2 < 2 || num2 > 36 {
		return -1
	}

	count := 0

	for {
		if num1 == 0 {
			break
		}
		if num1 != 0 {
			num1 = num1 / num2
			count++
		}
	}

	return count
}

func main() {
	fmt.Println(DigitLen(100, 10))
	fmt.Println(DigitLen(100, 2))
	fmt.Println(DigitLen(-100, 16))
	fmt.Println(DigitLen(100, -1))
}
