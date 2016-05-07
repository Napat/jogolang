/*
  - Concurrency/Multitasking: Doing many things but only one at a time(Single-core)
  - Parallelism: Doing many things at the same time(Multi-cores)
  - go routine: beware race condition when concurrency
  - wait-group
  - gomaxprocs_parallelism
    + Before go 1.5, default GOMAXPROCS is only one core
    + After go 1.5, default GOMAXPROCS is runtime.NumCPU()
  -	time sleep
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func foo() {
	for i := 0; i < 3; i++ {
		fmt.Println("Foo:", i)
		time.Sleep(300 * time.Millisecond)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 3; i++ {
		fmt.Println("Bar:", i)
		time.Sleep(500 * time.Millisecond)
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}
