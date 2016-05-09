package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, playground")
	go working()
	/*for {
		fmt.Println("main ticker")
		time.Sleep(1000 * time.Millisecond)
	}*/
	// playground demonstrate
	fmt.Println("main ticker 01")
	time.Sleep(10000 * time.Millisecond)
	fmt.Println("main ticker 02")
	time.Sleep(10000 * time.Millisecond)
	fmt.Println("main ticker 03")
	time.Sleep(10000 * time.Millisecond)
	fmt.Println("main ticker 04")
}
func working() {
	done := make(chan struct{}, 1)
	go a(done)
	go b(done)
	time.Sleep(20000 * time.Millisecond)
	// closing done then always immediately take <-done case in a() and b()
	close(done)
	fmt.Println("working done")
}
func a(done chan struct{}) {
	defer fmt.Println(".. a is ended ..")
	i := 1
	for {
		// blocking when receive done channel until closed
		// or receive time channel when after duration
		select {
		case <-done:
			fmt.Println("a done")
			// without return, select will infinite loop <-done case
			// Note: go playground will not allow you to try the loop
			return // remove return to see "a done" loop
		case <-time.After(1000 * time.Millisecond):
			fmt.Println("[A]", i)
			i++
		}
	}

}
func b(done chan struct{}) {
	defer fmt.Println(".. b is ended ..")
	i := 1
	for {
		// blocking when receive done channel until closed
		// or receive time channel when after duration
		select {
		case <-done:
			fmt.Println("b done")
			// without return, select will infinite loop <-done case
			// Note: go playground will not allow you to try the loop
			return // remove return to see "b done" loop
		case <-time.After(1500 * time.Millisecond):
			fmt.Println("[B]", i)
			i++
		}
	}
}
