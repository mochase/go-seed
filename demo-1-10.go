/*
显示的类型转换, 类型推导, 常量的关键字不能省略
*/

package main

import (
	"fmt"
	"math"
)

func main () {
	var x, y = 3, 4
	var f float64 = math.Sqrt(float64(x * x + y * y))
	var z = f
	const P = 3.14
	fmt.Println(x, y, z)
	fmt.Printf("v is of type %T\n", z)
	fmt.Println(P)
}