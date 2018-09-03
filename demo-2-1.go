/*
for 循环, 前置语句与后置语句都可选
循环条件也可以省略, 及无限循环
*/

package main

import "fmt"

func main() {
	sum := 0
	i := 0
	for ; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	fmt.Println(i)

	a := 1
	for a < 1000 {
		a += a
	}
	fmt.Println(a)
}