/*
map 结构
若顶级类型只是一个类型名，你可以在文法的元素中省略它
*/



package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["bell labs"] = Vertex{
		40.68433, -74.39967,
	}

	b := map[string]Vertex{
		"bell labs": Vertex{
			40.68433, -74.39967,
		},
		"google": {
			37.42202, -122.08408,
		},
	}

	b["tony"] = Vertex{
		123, 456,
	}

	delete(b, "google")
	v, ok := b["google"]

	fmt.Println(m["bell labs"])
	fmt.Println(b)
	fmt.Println("the value: ", v, "present?", ok)
}