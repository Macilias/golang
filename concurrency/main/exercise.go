package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	// send the value
	ch <- t.Value

	// walk the tree
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)
	go Walk(t1, c1)
	go Walk(t2, c2)

	// TODO complete

	return c1 == c2
}

func main() {

	fmt.Println("Should be true is", Same(tree.New(1), tree.New(1)))
	fmt.Println("Should be false is", Same(tree.New(1), tree.New(2)))

}
