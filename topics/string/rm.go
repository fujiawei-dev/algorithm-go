package main

import "fmt"

func main() {
	s := []byte("(1,(2,3),(4,(5,6),7))")
	fmt.Println(string(RemoveNested(s)))
}

func RemoveNested(s []byte) (ret []byte) {
	if s[0] != '(' || s[len(s)-1] != ')' {
		panic("invalid format")
	}
	ret = append(ret, '(')
	brackets := 0
	for i := range s {
		if s[i] == '(' {
			brackets++
		} else if s[i] == ')' {
			brackets--
		} else {
			ret = append(ret, s[i])
		}
	}
	if brackets != 0 {
		panic("invalid format")
	}
	return append(ret, ')')
}
