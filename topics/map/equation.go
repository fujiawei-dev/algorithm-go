package main

import "fmt"

func main() {
	list := []int{3, 4, 7, 10, 2, 9, 8}
	fmt.Println(FindEquation(list))
}

func FindEquation(list []int) ([2]int, [2]int) {
	kv := make(map[int][2]int)
	for idx, val := range list {
		for i := idx + 1; i < len(list); i++ {
			k := val + list[i]
			if v, ok := kv[k]; ok {
				return v, [2]int{val, list[i]}
			}
			kv[k] = [2]int{val, list[i]}
		}
	}
	return [2]int{0, 0}, [2]int{0, 0}
}
