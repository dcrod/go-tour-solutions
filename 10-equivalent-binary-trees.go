package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t (via trav) sending
// values from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	trav(t, ch)
	close(ch)
}

// trav does the actual tree traversal
func trav(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	trav(t.Left, ch)
	ch <- t.Value
	trav(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for v1 := range ch1 {
		v2, ok := <-ch2
		if !ok || v1 != v2 {
			return false
		}
	}
	return true
}

func main() {
	// Test Walk function
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for v := range ch {
		fmt.Println(v)
	}
	// Test the Same function
	fmt.Println(Same(tree.New(1), tree.New(1))) // true
	fmt.Println(Same(tree.New(1), tree.New(2))) // false
}
