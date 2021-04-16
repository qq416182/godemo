package main

import "fmt"

func main() {
	const LENGTH int = 10
	const WIDTH = 20
	const a, b, c = 3.4, false, "hello"
	area := LENGTH * WIDTH
	fmt.Println(area)
	fmt.Println(a, b, c)

}
