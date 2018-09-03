package main

import "fmt"

func main() {
	q := []int{2, 3, 4, 5, 6}
	r := []bool{true. false, true, false}
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{4, true},
		{7, true}
	}
	fmt.Println(q, r, s)
}