/*
  https://play.golang.org/p/78LH1I-l0i
  - Defer/ Panic / Recover
  - Get line executing function: fileLine()

	More panic sample
	https://golang.org/src/encoding/json/encode.go?s=5584:5627#L127
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

func f() (err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Found panic")
			if _, ok := r.(runtime.Error); ok {
				fmt.Println("runtime panic")
				panic(r)
			}
			if s, ok := r.(string); ok {
				fmt.Println(fileLine(), "found panic: ", s)
				panic(s)
			}
			err = r.(error) // forward error
		}
	}()
	g(0)
	fmt.Println("End of f()")
	return nil
}

type gError struct {
	count int
}

func (e *gError) Error() string {
	return fmt.Sprintf("Count=%v at %v", e.count, fileLine())
}

func g(count int) {
	defer fmt.Printf("Defer of g(%v)\n", count)
	fmt.Printf("Try g(%v)...", count)
	if count > 1 {
		fmt.Println("Panicking!")
		panic(&gError{count})
	}
	fmt.Printf("g(%v) is ok\n", count)
	g(count + 1)
}

func main() {
	if r := f(); r != nil {
		fmt.Println(r)
	}
	fmt.Println("End main()")
}
