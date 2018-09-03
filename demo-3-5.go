/*
切片
切片是数组的引用,本身并不存储数据,更改切片的元素会修改底层的数组
*/

package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[:4]
	var b []int = primes[1:5]
	s[2] = 99
	fmt.Println(s, b)
	fmt.Println(primes)
}