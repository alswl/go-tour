package main

import "fmt"

// builtin https://golang.org/pkg/builtin/
/**
append, for slice
cap, array / pointer / slice / channel
close, channel
complex, number
copy, slice / bytes
delete, map
imag, imaginary of complex
len, array / pointer / slice, string, channel
make, slice / map / chan
new, allocate memory
panic,
print
println
real, read of complex
recover
**/

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {

	var s []int
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	s = append(s, 2, 3, 4)
	printSlice(s)

}
