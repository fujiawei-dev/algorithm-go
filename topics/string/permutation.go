package main

import "fmt"

func main() {
	s := []rune("abc")
	Permutation(s, 0)
	fmt.Println(result)
	fmt.Println(len(result))
}

var result []string

// Permutation 对字符串中的字符进行全排列
func Permutation(s []rune, start int) {
	if s == nil {
		return
	}
	if start == len(s)-1 {
		result = append(result, string(s))
	} else {
		for i := start; i < len(s); i++ {
			s[start], s[i] = s[i], s[start]
			Permutation(s, start+1)
			s[start], s[i] = s[i], s[start]
		}
	}
}
