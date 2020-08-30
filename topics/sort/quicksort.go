package main

import "fmt"

func main() {
	array := []int{5, 2, 3, 1, 9, 6, 4, 8, 7}
	QuickSort(array, 0, len(array)-1)
	fmt.Println(array)
}

// Partition 用于快速排序中的分割
func Partition(array []int, low, high int) int {
	pos, val := low, array[low]
	low++
	for low <= high {
		if array[low] < val {
			array[low], array[pos] = array[pos], array[low]
			low++
			pos++
		} else {
			array[high], array[low] = array[low], array[high]
			high--
		}
	}
	array[pos] = val
	return pos
}

// QuickSort 快速排序
func QuickSort(array []int, low, high int) {
	if low > high {
		return
	}
	pos := Partition(array, low, high)
	QuickSort(array, low, pos-1)
	QuickSort(array, pos+1, high)
}
