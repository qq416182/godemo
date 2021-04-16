package main

import "fmt"

func main() {
	var num int
	fmt.Scanf("%d", &num)
	num = 10
	if num > 10 {
		fmt.Println("the num is >10")
	} else if num == 10 {
		fmt.Println("the num is equal 10")
	} else {
		fmt.Println("the num is <10")
	}
}
