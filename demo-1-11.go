package main

import "fmt"

const (
	Big = 1 << 62
	Small = Big >> 63
)

func needInt (x int) int {
	return  x * 10 + 1
}

func needFloat (x float64) float64 {
	return x * 0.1
}

func main () {
	fmt.Println(Big)
	fmt.Println(Small)
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}