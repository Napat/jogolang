/*
Basic receiver and pointer of receiver
http://play.golang.org/p/Cck6WHSKwf
*/
package main

import "fmt"

type Counter struct {
	count int
}

func (c Counter) currentValue() int {
	return c.count
}
func (c Counter) notIncrement() {
	c.count++
}
func (c *Counter) increment() {
	c.count++
}

func main() {
	counter := Counter{1}

	counter.notIncrement()
	fmt.Printf("current value %d\n", counter.currentValue()) //current value 1

	counter.increment()
	fmt.Printf("current value %d\n", counter.currentValue()) //current value 2
}
