package main

func RetainFirstHalf(str string) string {

	if len(str) == 1 {
		return str
	}
	if len(str) == 0 {
		return ""
	}
	return str[:len(str)/2]
}
