package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var x int = 0 //共享变量
var rwlock sync.RWMutex

func main() {
	wg.Add(10)
	//启动7个reader
	for i := 0; i < 7; i++ {
		go reader(i)
	}
	//启动3个writer
	for i := 0; i < 3; i++ {
		go writer(i)
	}
	wg.Wait()
}

func reader(num int) {
	for {
		rwlock.RLock() //读锁
		fmt.Printf("i am %d reader goroutine x=%d\n", num, x)
		time.Sleep(time.Microsecond * 2)
		rwlock.RUnlock() //解锁
	}
	wg.Done()
}

func writer(num int) {
	for {
		rwlock.Lock() //读锁
		x += 1
		fmt.Printf("i am %d writer goroutine x=%d\n", num, x)
		time.Sleep(time.Microsecond * 2)
		rwlock.Unlock() //解锁
	}
	wg.Done()
}
