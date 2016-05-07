/*
Race-condition sample, just try
- go run -race main.go
- go run main.go
*/
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int //global share variable

func main() {
	wg.Add(2)
	go incrementor("Foo:")
	go incrementor("Bar:")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incrementor(s string) {
	for i := 0; i < 3; i++ {
		// Shared variable(without lock) will be cause of rase-condition
		counter++
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		fmt.Println(s, i, "Counter:", counter)
	}
	wg.Done()
}
