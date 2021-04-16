package main

import (
	"fmt"
	"sync"
)

var x int = 0
var wg sync.WaitGroup
var mutex sync.Mutex

func increment() {
	mutex.Lock()
	x += 1
	mutex.Unlock()
	wg.Done()
}
func main() {
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment()
	}
	wg.Wait()
	fmt.Println(x)
}
