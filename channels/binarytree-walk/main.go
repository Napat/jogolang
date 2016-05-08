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

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	/*
		*** Note ***
		To recursion in anonymous function need to
		declared function variable first.
		So, cannot self-executing
	*/
	var walker func(*tree.Tree)
	walker = func(tr *tree.Tree) {
		for tr != nil {
			walker(tr.Left)
			ch <- tr.Value
			tr = tr.Right
		}
	}
	walker(t)
	close(ch)
	return
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)

	go Walk(t1, c1)
	go Walk(t2, c2)

	for v1 := range c1 {
		v2, ok := <-c2
		//ok for checking btree dimension
		if v1 != v2 || !ok {
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

	for v := range ch {
		fmt.Print(v, " ")
	}
	fmt.Println()

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
