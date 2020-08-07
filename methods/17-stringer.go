package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Jingchao Di", 31}
	z := Person{"Jiuqing Di", 2}
	fmt.Println(a, z)
}
