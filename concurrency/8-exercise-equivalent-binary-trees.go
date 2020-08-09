package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

func contains(ts []*tree.Tree, ele *tree.Tree) bool {
	for _, a := range ts {
		if a == ele {
			return true
		}
	}
	return false
}

func Walk(t *tree.Tree, ch chan int) {
	recurse(t, ch)
	close(ch)
}

func recurse(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	recurse(t.Left, ch)
	ch <- t.Value
	recurse(t.Right, ch)
}

func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)

	go Walk(t1, c1)
	go Walk(t2, c2)
	for v := range c1 {
		if v != <-c2 {

			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for v := range ch {
		fmt.Println(v)
	}

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))

}
