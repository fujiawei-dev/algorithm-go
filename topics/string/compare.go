package main

import "fmt"

func main() {
	s1 := "aaaabbc"
	s2 := "abcbaaa"
	fmt.Println(Compare(s1, s2))
	fmt.Println(Contain(s1, s2))
}

// Contain 判断两个字符串的包含关系
func Contain(s1, s2 string) bool {
	if len(s1) < len(s2) {
		return false
	}
	dict := make(map[uint8]int)
	for i := range s1 {
		dict[s1[i]] += 1
		dict[s2[i]] -= 1
	}
	for i := range dict {
		if dict[i] < 0 {
			return false
		}
	}
	return true
}

// Compare 判断组成字符串的字符是否相同
func Compare(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	dict := make(map[uint8]int)
	for i := range s1 {
		dict[s1[i]] += 1
		dict[s2[i]] -= 1
	}
	for i := range dict {
		if dict[i] != 0 {
			return false
		}
	}
	return true
}
