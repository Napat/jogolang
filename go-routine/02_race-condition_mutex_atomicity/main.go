/*
- Race-condition checking
	+ go run -race main.go
- mutex (My preferred)
- atomicity
	+ import("sync/atomic")
*/
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int64 //global share variable
var mutex sync.Mutex

func main() {
	wg.Add(2)
	go incrementor("Foo:")
	go incrementor("Bar:")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incrementor(s string) {
	for i := 0; i < 3; i++ {

		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)

		/*//1. Shared variable(without lock) will cause rase-condition
		counter++
		fmt.Println(s, i, "Counter:", counter)
		*/

		//2. Mutex lock to solve race-condition
		mutex.Lock()
		counter++
		fmt.Println(s, i, "Counter:", counter)
		mutex.Unlock()

		/* // 3. Atomicity
		//Caution: Println still cause race-condition because shared variable
		atomic.AddInt64(&counter, 1)
		fmt.Println(s, i, "Counter:", counter)
		*/
	}
	wg.Done()
}
