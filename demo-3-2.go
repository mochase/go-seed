/*
struct结构体
结构体指针: (*p).X,简略写法: p.X
*/

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 4
	p := &v
	p.Y = 1e9
	fmt.Println(v)
}