/*
  https://play.golang.org/p/78LH1I-l0i
  -Defer/ Panic / Recover
  -get line executing function: fileLine()
*/

package main

import "fmt"
import "runtime"

func fileLine() string {
	var s string
	_, fileName, fileLine, ok := runtime.Caller(1)
	/*Pass 1 to Caller() returns the number of the line where fileLine() is called*/

	if ok {
		s = fmt.Sprintf("%s:%d", fileName, fileLine)
	} else {
		s = ""
	}
	return s
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Found panic: ", r)
			fmt.Println("Recovered")
		}
	}()
	g(0)
	fmt.Println("End of f()")
}

func g(count int) {
	defer fmt.Printf("Defer of g(%v)\n", count)
	fmt.Printf("Try g(%v)...", count)
	if count > 1 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("Count=%v at %v", count, fileLine()))
	}
	fmt.Printf("g(%v) is ok\n", count)
	g(count + 1)
}

func main() {
	f()
	fmt.Println("End main()")
}
