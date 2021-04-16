package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("i am main goroutine")
	//一个goroutine
	go func() {
		fmt.Println("i am a goroutine")
	}()
	fmt.Println("byebye")
	time.Sleep(time.Second * 2)
}
