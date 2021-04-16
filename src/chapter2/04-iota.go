package main

import "fmt"

const (
	login = iota
	logout
	user    = iota + 2 //2+2
	account = iota * 3 //9
)

const (
	apple, banana = iota + 1, iota + 2
	peach, pear
	orange, mango
)

func main() {
	fmt.Println(login, logout, user, account)
	fmt.Println(apple, banana, peach, pear, orange, mango)
}
