/*
if 语句,可以在条件表达式前执行简单的赋值语句;
块级作用域
if 块中的变量可以被 else 块捕获
*/

package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v 
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main () {
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(pow(3, 2, 10))
	fmt.Println(pow(3, 3, 10))
}