package main

import "fmt"

func main() {
	nextNumber := getSequece()
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
}

func getSequece() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
