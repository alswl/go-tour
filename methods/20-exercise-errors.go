package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	// value e, cause cascade Error(), use float64(e)
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	var guess, accurate = 1.0, 0.000000001
	for ; math.Abs(x-guess*guess) > accurate; guess -= (guess*guess - x) / (2 * guess) {
		//fmt.Printf("guess: %g\n", guess)
	}

	return guess, nil
}

func main() {

	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
