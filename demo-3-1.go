/*
指针运算 
*/

package main

import "fmt"

func main() {
	i, j := 42, 2701
	var p *int = &i
	fmt.Println(*p)
	fmt.Println(p)
	*p = 21
	fmt.Println(i)
	p = &j
	*p = *p / 37
	fmt.Println(j)
}