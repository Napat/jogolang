package main

import "fmt"

func main() {

// Slices hold references to an underlying array, and if you assign one slice to another, both refer to the same array.

	//s1 := []int{1,2,3,4,5}
	s1 := make([]int, 5,5); s1[0] = 1;s1[1] = 2;s1[2] = 3;s1[3] = 4;s1[4] = 5;
	
	fmt.Println(s1)
	s2 := s1
	s1[0] = 0
	fmt.Println(s1[0] ,s2[0])
	fmt.Println(cap(s1) ,cap(s2))
	s1 = append(s1 ,6)
	fmt.Println(cap(s1) ,cap(s2))

/* Output
[1 2 3 4 5]
0 0
5 5
12 5
*/

// If a function takes a slice argument, changing elements of the slice inside function will be visible to the caller,
// analogous to passing a pointer to the underlying array
	s3 := []int{1,2,3,4,5}
	fmt.Println(s3)
	likePointerArray(s3)
	fmt.Println(s3)
/* Output
[1 2 3 4 5]
[0 2 3 4 5]
*/

}

func likePointerArray(s []int){
	s[0] = 0
}
