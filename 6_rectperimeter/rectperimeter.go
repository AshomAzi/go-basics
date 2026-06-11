package main

import "fmt"

func rectperimeter(width, height int) int {

	if width < 0 || height < 0 {
		return -1
	}

	return 2*(width+height)
}

func main() {
	fmt.Println(rectperimeter(2,2))
}