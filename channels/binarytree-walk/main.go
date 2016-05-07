/*
Solve of
  https://tour.golang.org/concurrency/7
  https://tour.golang.org/concurrency/8
*/
package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

const bTreeLen = 10

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
	return
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)

	go Walk(t1, c1)
	go Walk(t2, c2)

	var v1, v2 int
	for i := 0; i < bTreeLen; i++ {
		v1, v2 = <-c1, <-c2
		if v1 != v2 {
			return false
		}
	}
	return true
}

func main() {
	t1 := tree.New(1)
	fmt.Println(t1.String())

	var x *tree.Tree
	fmt.Println(x.String())

	ch := make(chan int)
	go Walk(t1, ch)

	for i := 0; i < bTreeLen; i++ {
		fmt.Print(<-ch, " ")
	}
	fmt.Println()

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
