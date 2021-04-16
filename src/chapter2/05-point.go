package main

import "fmt"

func main() {
	a := 10
	var b *int = &a
	fmt.Println(*b, b)
	*b = 100
	fmt.Println(a, *b, b)
}
