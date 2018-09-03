/*
bool
string
int, int8, int16, int32, int64
uint, uint8, uint16, uint32, uint64, uintptr
byte (uint8)
rune (int32, 表示一个 Unicode 码点)
float32, float64
complex64, complex128
*/

package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe bool = false
	MaxInt uint64 = 1<<3 - 1
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

