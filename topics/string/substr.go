package main

import (
	. "algorithm-go/structs"
	"fmt"
)

func main() {
	s1 := []rune("1KEPewhjGZ6ApcgyKtksZEPewhjGZ6Ap")
	s2 := []rune("IYhwhjGZmfjwgyKtksZmmF2KtksZEPe")

	fmt.Println(string(LongestSubString(s1, s2)))
	fmt.Println(string(LongestSubStringDyn(s1, s2)))
}

//BenchmarkLongestSubStringDyn-8            299994              3650 ns/op            9344 B/op         34 allocs/op
//BenchmarkLongestSubString-8               630778              1994 ns/op               0 B/op          0 allocs/op

// LongestSubStringDyn 动态规划法求两个字符串的最长公共子串，性能略差
func LongestSubStringDyn(s1, s2 []rune) (sub []rune) {
	l1, l2 := len(s1), len(s2)
	pos, max := 0, 0                 // 用来记录最长公共字串开始和结束字符的位置
	matrix := InitMatrix(l1+1, l2+1) // 申请新的空间来记录公共字串长度信息
	for i := 1; i < l1+1; i++ {
		for j := 1; j < l2+1; j++ {
			if s1[i-1] == s2[j-1] {
				matrix[i][j] = matrix[i-1][j-1] + 1
				if matrix[i][j] > pos {
					pos = matrix[i][j]
					max = i
				}
			} else {
				matrix[i][j] = 0
			}
		}
	}
	//PrintMatrix(matrix)
	return s1[max-pos : max]
}

// LongestSubString 求两个字符串的最长公共子串
func LongestSubString(s1, s2 []rune) (sub []rune) {
	p, q := 0, 0 // 记录最长子序列的开始和结束
	for k := 0; k < len(s2); k++ {
		i, j, x, y := 0, 0, 0, k         // 记录子序列的开始和结束
		for x < len(s1) && y < len(s2) { // 字符序列当前下标
			if s1[x] != s2[y] {
				if i != j { // 说明之前有过相同元素
					y = k // 必须重置起始位置
				} else {
					x++
				}
				i, j = x, x
			} else {
				j, x, y = j+1, x+1, y+1
			}
			// 判断当前子序列是否更长
			if j-i > q-p {
				q, p = j, i
			}
		}
	}
	return s1[p:q]
}
