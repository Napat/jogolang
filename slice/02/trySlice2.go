package main

import "fmt"

func main() {
	numArr := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%T %d %d ", numArr, len(numArr), cap(numArr))
	fmt.Println(numArr)
	fmt.Println("----------------")


	numSlice := numArr[2:3] // cap count from [idx=2] to [endIdx=3] = 3 
	fmt.Printf("%T %d %d ", numSlice, len(numSlice), cap(numSlice))
	fmt.Println(numSlice)
	fmt.Println("----------------")

	for i := 3; i <= 6; i++ {
		// cap count from [idx=2] to [idx=i]
		// cap must greater or equal numbers of element in new slice 
		// and must not out of old array bound
		numSlice2 := numArr[2:3:i]	 	
		fmt.Printf("%T %d %d ", numSlice2, len(numSlice2), cap(numSlice2))
		fmt.Println(numSlice2)
	}
}

// Go Slicing from Array
// https://www.facebook.com/iporsut/media_set?set=a.10154064633924924.1073741834.640039923&type=3

/* Output

[5]int 5 5 [1 2 3 4 5]
----------------
[]int 1 3 [3]
----------------
[]int 1 1 [3]
[]int 1 2 [3]
[]int 1 3 [3]
panic: runtime error: slice bounds out of range

*/
