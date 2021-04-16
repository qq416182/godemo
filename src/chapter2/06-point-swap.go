package main

import "fmt"

func main() {
	a, b := 10, 20
	swap(a, b)
	fmt.Println(a, b)
	swap2(&a, &b)
	fmt.Println(a, b)
}

func swap(a, b int) {
	temp := a
	a = b
	b = temp
}

func swap2(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
