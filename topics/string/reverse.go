package main

import "fmt"

func main() {
	s := []rune("how about you")

	ReverseString(s)
	start := 0
	for i := range s {
		if s[i] == ' ' {
			ReverseString(s[start:i])
			start = i + 1
		}
	}
	ReverseString(s[start:])

	fmt.Println(string(s))
}

func ReverseString(s []rune) []rune {
	front, rear := 0, len(s)-1
	for front < rear {
		s[front] = s[rear] ^ s[front]
		s[rear] = s[rear] ^ s[front]
		s[front] = s[rear] ^ s[front]
		front++
		rear--
	}
	return s
}
