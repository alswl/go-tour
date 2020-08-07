package main

import (
	"fmt"
	"math"
)

func powWithLim2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
	fmt.Println(
		powWithLim2(3, 2, 10),
		powWithLim2(3, 3, 20),
	)
}
