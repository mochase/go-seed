/*
零值, 没有 undefined
*/

package main

import "fmt"

func main () {
	var (
		i int
		f float64
		b bool
		s string
	)
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}