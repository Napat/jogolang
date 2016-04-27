package main

import "fmt"

func main() {
	// Createing slice
	//nums := []int{}			// shorthand method
	//var nums []int			// var method: not initialized array then nil
	nums := make([]int, 2, 4)	// make method
	//nums := new([4]int)[0:2]	// new and slicing method(same as make method)

	fmt.Println(len(nums))
	fmt.Println(cap(nums))
	fmt.Println(nums)
	fmt.Println(nums == nil)
	fmt.Printf("%p\n", nums)
	fmt.Println("-----")

	for i := 0; i < 20; i++ {
		nums = append(nums, i+3)
		fmt.Println(len(nums))
		fmt.Println(cap(nums))
		fmt.Println(nums)
		fmt.Println("-----")
	}

}
