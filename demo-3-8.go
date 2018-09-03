/*
切片 的 "foreach"循环
*/

package main

import (
	"fmt"
	"strings"
)


func main() {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for _, v := range board {
		fmt.Printf("%s\n", strings.Join(v, " "))
	}
}