/*
切片,
长度和容量,长度为切片的上标-下标, 容量为底层数组的上标-切片的下标
注意切片的默认值
用 make 创建切片
*/

package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	s = s[:3]
	printSlice(s)

	s = s[1:]
	printSlice(s)

	s = s[2:4]
	printSlice(s)

	a := make([]int, 5)
	printSlice(a)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

/*
len=6 cap=6 [2 3 5 7 11 13]
len=3 cap=6 [2 3 5]
len=2 cap=5 [3 5]
len=2 cap=3 [7 11]
*/