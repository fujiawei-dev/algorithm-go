---
date: 2020-10-12T17:08:42+08:00  # 创建日期
author: "Rustle Karl"  # 作者

# 文章
title: "基本数学运算"  # 文章标题
description: "位运算、对数运算等"
url:  "posts/algorithm/go/docs/math"  # 设置网页永久链接
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

## 位运算

### 位运算符号

| 符号 | 含义 |
| ---- | ---- |
| & | AND |
| \| | OR |
| ^ | XOR / 取反 |
| &^ | 位清空 |
| << | 左移 |
| >> | 右移 |

### 数的二进制表示

```go
func main() {
	n := 3
	fmt.Printf("%08b\n", uint8(n))
	fmt.Printf("%08b\n", uint8(-n))  // 取反 +1
	//fmt.Printf("%b\n", uint8(-3))
	// Cannot convert expression of type 'int' to type 'uint8'
}
```

```
00000011
11111101
```

> 无法将负数直接转换为 `uint` 类型，但是赋值给一个变量之后就可以了，暂时不知道为什么。

### 判断二进制数某位是否为 1

```go
// 判断二进制数从右往左数第 i 位是否为 1，位数从 0 开始
func IsOneInt(n, i int) bool {
	return (1<<i)&n == 1<<i
}

// 从结果上看两者没有区别
func IsOneUint(n, i int) bool {
	return (uint(n) & (uint(1) << uint(i))) == uint(1)<<uint(i)
}
```

### 求绝对值

仅用于学习，实际性能与取负数是差不多的。

```go
func Abs(n int) int {
	if n >= 0 {
		return n
	}
	return ^(n - 1)
}
```

## 对数运算

```go
func main() {
    // 以自然对数 e 为底
    fmt.Println(math.Log(10))
    // 以 2 为底
    fmt.Println(math.Log2(16))  // 4
    // 以 10 为底
	fmt.Println(math.Log10(100))  // 2
    // 对数运算法则
	fmt.Println(math.Log(16)/math.Log(2))  // 2
}
```
