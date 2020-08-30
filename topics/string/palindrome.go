package main

import "fmt"

func main() {
	s := "cabcba"
	fmt.Println(IsPalindrome(s))
}

func findLongestPal(s string) string {

}

func IsPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
