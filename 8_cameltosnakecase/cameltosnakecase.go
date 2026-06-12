package main

import (
	"fmt"
)

func CamelToSnakeCase(str string) string {
	if str == "" {
		return ""
	}

	for i := 0; i < len(str); i++ {
		c := str[i]
		if !((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')) {
			return str
		}
		if c >= 'A' && c <= 'Z' && i > 0 {
			prev := str[i-1]
			if prev >= 'A' && prev <= 'Z' {
				return str
			}
		}
	}

	last := str[len(str)-1]
	if last >= 'A' && last <= 'Z' {
		return str
	}

	result := ""
	for i := 0; i < len(str); i++ {
		c := str[i]
		if c >= 'A' && c <= 'Z' {
			if i > 0 {
				result += "_"
			}
			result += string(c + ('a' - 'A'))
			continue
		}
		result += string(c)
	}

	return result
}

func main() {
	fmt.Print(CamelToSnakeCase("WelcomdBeackHome"))
}
