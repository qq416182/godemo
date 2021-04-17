package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 100; i++ {
		go func(num int) {
			time.Sleep(time.Second * time.Duration(num))
			fmt.Println("go", num)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("-----------------")
}
