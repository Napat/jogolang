package main

import "fmt"

/*
https://golang.org/doc/effective_go.html#arrays

There are major differences between the ways arrays work in Go and C. In Go,

Arrays are values. Assigning one array to another copies all the elements.
In particular, if you pass an array to a function, it will receive a copy of the array, not a pointer to it.
The size of an array is part of its type. The types [10]int and [20]int are distinct.

***The value property can be useful but also expensive;
if you want C-like behavior and efficiency,
you can pass a pointer to the array.***
*/

func main() {
	v1 := [5]int{1, 2, 3, 4, 5}                    // basic reference type array
	fmt.Printf("%T\t %T\t %d\n", v1, v1[1], v1[1]) // access array value

	v2 := &[5]int{1, 2, 3, 4, 5}                           // pointer to the array
	fmt.Printf("%T\t %T\t %d\n", v2, v2[1], v2[1])         // !! Weird!! Can access value without indirect access 
	fmt.Printf("%T\t %T\t %d\n\n", v2, (*v2)[1], (*v2)[1]) // Using indirect access still working

	a1 := [5]int{1, 2, 3, 4, 5}
	a2 := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%d %d\n", a1[1], a2[1])
	unLikeC(a1) 
	toLikeC(&a2)
	fmt.Printf("%d %d\n", a1[1], a2[1])
}

// The argument values will copying to array parameter(Not reference type like C pointer)
// Slow speed but can modify parameter inside function
func unLikeC(a1 [5]int) { 
	a1[1] = 10
}

// pass a pointer argument to the array parameter (Reference type)
func toLikeC(a2 *[5]int) {
	a2[1] = 10
}
