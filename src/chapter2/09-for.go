package main

import (
	"fmt"
	"time"
)

func main() {
	var i, sum int = 0, 0
	for i = 1; i <= 2000; i++ {
		sum += i
	}
	fmt.Println(sum)
	i = 1
	sum = 0
	for i <= 1000 {
		sum += i
		i++
	}
	fmt.Println(sum)

	for {
		time.Sleep(time.Second * 1)
		fmt.Println("hehe")
	}
}
