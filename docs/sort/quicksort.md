---
date: 2020-11-20T11:03:48+08:00  # 创建日期
author: "Rustle Karl"  # 作者

# 文章
title: "快速排序 Quick Sort"  # 文章标题
url:  "posts/algorithm/go/docs/sort/quicksort"  # 设置网页永久链接
tags: [ "algorithm", "go" ]  # 标签
series: [ "Go 数据结构与算法"]  # 系列
categories: [ "数据结构与算法"]  # 分类

# 章节
weight: 20 # 排序优先级
chapter: false  # 设置为章节

index: true  # 是否可以被索引
toc: true  # 是否自动生成目录
draft: false  # 草稿
---

## 算法实现

```go
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
```

## 基于快速选择的中位数查找

```go
func QuickSelectMedian(seq []float64, low int, high int, k int) float64 {
	if low == high {
		return seq[k]
	}
	for low < high {
		pivot := low/2 + high/2
		pivotValue := seq[pivot]
		storeIdx := low
		seq[pivot], seq[high] = seq[high], seq[pivot]
		for i := low; i < high; i++ {
			if seq[i] < pivotValue {
				seq[storeIdx], seq[i] = seq[i], seq[storeIdx]
				storeIdx++
			}
		}
		seq[high], seq[storeIdx] = seq[storeIdx], seq[high]
		if k <= storeIdx {
			high = storeIdx
		} else {
			low = storeIdx + 1
		}
	}
	if len(seq)%2 == 0 {
		return seq[k-1]/2 + seq[k]/2
	}
	return seq[k]
}
```

```go

```

```go

```

```go

```

