package main

import (
	"fmt"
)

func factor(z, x float64) float64 {
	return z - (z * z - x) / (2 * z);
}

func factor_wrapper(z float64, x float64, count int) float64 {
	if count > 100 {
		return factor(z, x)
	}
	count++
	return factor_wrapper(factor(z, x), x, count)
}

func Sqrt(x float64) float64 {
	z := 5.0

	return factor_wrapper(z, x, 0)
}

func main() {
	fmt.Println(Sqrt(121))
}
