--Unbuffered channel--
ch := make(chan int)
By default, sends and receives block until the other side is ready.
This allows goroutines to synchronize without explicit locks or condition variables.

--Buffered Channels--
ch := make(chan int, 100)
Channels can be buffered. Provide the buffer length as the second argument
to make to initialize a buffered channel.
Sends to a buffered channel block only when the buffer is full.
Receives block when the buffer is empty.

--Sender--
can close a channel to indicate that no more values will be sent,
close(c)

--Receivers--
can test whether a channel has been closed by assigning a
second parameter to the receive expression,
v, ok := <-ch
ok is false if there are no more values to receive and the channel is closed.

--Range--
The loop for i := range c receives values from the channel repeatedly until it is closed.

--Note--
- Only the sender should close a channel, never the receiver.
Sending on a closed channel will cause a panic.
- Channels aren't like files; ***you don't usually need to close them.***
Closing is only necessary when the receiver must be told there are no more values coming,
such as to terminate a range loop

--Select--
The select statement lets a goroutine wait on multiple communication operations.

- select without default
A select blocks until one of its cases can run, then it executes that case.
It chooses one at random if multiple are ready.

https://tour.golang.org/concurrency/5
package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
--------------------------------

- select with default
package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(1000 * time.Millisecond)
	boom := time.After(5000 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
-------------------------------------



https://tour.golang.org/concurrency/8


package main

import "fmt"
import "golang.org/x/tour/tree"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int){
	ch <- t.Value
	if t.Left != nil {
		go Walk(t.Left,ch)
	}
	if t.Right != nil {
		go Walk(t.Right,ch)
	}
}
/*
// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool
*/
func main() {
	t1:= tree.New(1)
	fmt.Println(t1.String())

	ch := make(chan int)
	go Walk(t1, ch)

	for i := range ch {
		fmt.Println(i)
	}

}
