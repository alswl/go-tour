package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Split(s, " ")
	counts := make(map[string]int, len(words))

	for _, v := range words {
		if _, ok := counts[v]; ok {
			counts[v] += 1
		} else {
			counts[v] = 1
		}
	}
	
	return counts
}

func main() {
	wc.Test(WordCount)
}
