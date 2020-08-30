package main

import "fmt"

func main() {

	s := []byte("AbcDef")
	fmt.Println(string(sortLetter(s)))
}

func sortLetter(s []byte) []byte {
	front, rear := 0, len(s)-1
	for front < rear {
		for s[front] >= 'a' { // A 65 / a 97
			front++
		}
		for s[rear] < 'a' {
			rear--
		}
		if front < rear {
			s[front], s[rear] = s[rear], s[front]
		}
	}
	return s
}
