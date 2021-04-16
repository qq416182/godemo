package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(time.Microsecond * 100)
	fmt.Println("\n", fib(45))
}

func fib(x int) int {
	//f(x) = f(x-1) + f(x+2)  x=0,f=1,x=1,f=1
	if x < 2 {
		return x
	}
	return fib(x-2) + fib(x-1)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range "-\\|/" {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}
