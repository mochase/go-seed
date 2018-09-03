package main

import (
	"fmt"
)

func fibonacci() func() int {
	a := 0
	b := 1
	return func() int {
		c := b
		b = a + b
		a = c
		return a
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i ++ {
		fmt.Println(f())
	}
}
