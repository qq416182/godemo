package main

import (
	"fmt"
)

func main() {
	var v1 uint
	var x, y = 123, "hello"
	u, t := 456, "world"
	fmt.Println(v1, x, y, u, t)
	var cl complex64 = 3 + 4i
	fmt.Println(cl)
}
