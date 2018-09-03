/*
defer
defer 栈,
defer 函数的参数已经被绑定了
*/

package main

import "fmt"

func main() {
	fmt.Println("counting")
	defer fmt.Println("done")
	for i :=0; i < 10; i++ {
		defer fmt.Println(i)
	}
}