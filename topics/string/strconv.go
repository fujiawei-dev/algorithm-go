package main

import "fmt"

func main() {
	fmt.Println(IsInteger("-123"))
	fmt.Println(IsInteger("+123"))
	fmt.Println(IsInteger("123"))
	fmt.Println(IsInteger("-+123"))
}

func IsInteger(s string) (int, bool) {
	if s == "" {
		return 0, false
	}
	neg := 1
	if s[0] == '+' {
		s = s[1:]
	} else if s[0] == '-' {
		s = s[1:]
		neg = -1
	}
	var n uint8
	for i := range s {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}
		n = n*10 + s[i] - '0'
	}
	return neg * int(n), true
}
