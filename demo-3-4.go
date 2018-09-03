/*
数组
不能改变长度
*/
package main

/*
数组, 长度固定
*/
import "fmt"

func main() {
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 4, 11, 13}
	fmt.Println(primes)
}