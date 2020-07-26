package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	(*p).X = 20
	fmt.Println(v.X)
	// alias (*p).X
	p.X = 10
	fmt.Println(v.X)
}