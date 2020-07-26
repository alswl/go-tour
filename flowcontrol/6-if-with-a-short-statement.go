package main

import (
	"fmt"
	"math"
)

func powWithLim(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func pow(x, n float64) float64 {
	v := math.Pow(x, n)
	return v
}

func main() {
	fmt.Println(
		powWithLim(3, 2, 10),
		powWithLim(3, 3, 20),
	)

	fmt.Println(
		pow(3, 2),
		pow(3, 3),
	)

}
