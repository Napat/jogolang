/*
	- Init function
	- Defer behavior
*/

package main

import "fmt"

func defer01() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

func defer02() {
	defer fmt.Print("\n")
	for i := 0; i < 4; i++ {
		defer fmt.Print(i)
	}
}

func defer03() (ret int) {
	//defer func() { ret++ }()
	return 1
}

func init() {
	fmt.Println("Go automatic calling init function")
}

func main() {
	defer01()              // 0
	defer02()              // 3210
	fmt.Println(defer03()) // 2
}
