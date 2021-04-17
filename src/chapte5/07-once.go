package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int = 0 //共享数据
	var wg sync.WaitGroup
	var once sync.Once

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			once.Do(func() {
				count += 1 //只会执行一次，ONCE的作用
			})
			wg.Done()
		}()
	}
	wg.Wait() //阻塞等待所有goroutine结束
	fmt.Println(count)
}
