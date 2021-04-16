package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("max", runtime.GOMAXPROCS(1))
	runtime.GOMAXPROCS(3) //设置
	fmt.Println("max", runtime.GOMAXPROCS(1))
	//runtime.Goexit()
	runtime.Gosched() //释放CPU
}
