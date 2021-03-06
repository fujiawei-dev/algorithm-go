package main

import "fmt"

func main() {
	array := []int{5, 2, 3, 1, 9, 6, 4, 8, 7}
	fmt.Println(findTop3(array))
}

func findTop3(array []int) (r1, r2, r3 int) {
	if array == nil || len(array) < 3 {
		return
	}
	for _, v := range array {
		if v > r1 {
			r3 = r2
			r2 = r1
			r1 = v
		} else if v > r2 && v != r1 {
			r3 = r2
			r2 = v
		} else if v > r3 && v != r2 {
			r3 = v
		}
	}
	return
}
