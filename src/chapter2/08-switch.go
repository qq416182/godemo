package main

import "fmt"

func main() {
	var fruit string
	fmt.Scanf("%s", &fruit)
	switch fruit {
	case "apple":
		fmt.Println("this is an apple.")
	case "orange":
		fmt.Println("this is an orange.")
	case "pear":
		fmt.Println("this is a pear.")
	default:
		fmt.Println("are you kidding me?")
	}
}
