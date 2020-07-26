package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var guess, accurate = 1.0, 0.000000001

	for ; math.Abs(x-guess*guess) > accurate; guess -= (guess*guess - x) / (2 * guess) {
		fmt.Printf("guess: %g\n", guess)
	}

	return guess
}

func main() {
	fmt.Println(Sqrt(2))

}
