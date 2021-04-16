package main

import "fmt"

func main() {
	fmt.Println(add_sub(4, 1))
	fmt.Println(action(100, 10, add_action))
	fmt.Println(action(100, 10, sub_acction))
}

func add_sub(i, s int) (int, int) {
	return i + s, i - s
}

func add_action(a, b int) int {
	return a + b
}

func sub_acction(a, b int) int {
	return a - b
}

func action(a, b int, math func(a, b int) int) int {
	return math(a, b)
}
