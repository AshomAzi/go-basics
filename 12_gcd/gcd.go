package main

import "fmt"

func Gcd(num1, num2 uint) int {
	if num1 == 0 || num2 == 0 {
		return 0
	}
	num3 := uint(0)

	for {

		num3 = num1%num2 //12
		num1 = num2
		num2 = num3
		if num2 == 0{
			break
		}
	}

	return int(num1)
}

// func Gcd(num1, num2 uint) int {
// 	if num1 == 0 || num2 == 0 {
// 		return 0
// 	}
// 	num3 := uint(0)

// 	for {

// 		num3 = num1 % num2 //12
// 		num1 = num2

// 		if num3 != 0 {
// 			num2 = num3

// 		} else {
// 			break
// 		}
// 	}

// 	return int(num1)
// }



func main() {
	fmt.Println(Gcd(42, 10))
	fmt.Println(Gcd(42, 12))
	fmt.Println(Gcd(14, 77))
	fmt.Println(Gcd(17, 3))
}
