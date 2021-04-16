package main

import "fmt"

type Person struct {
	name  string
	age   int
	sex   string
	fight int
}

func main() {
	p := Person{
		name:  "洒洒",
		age:   20,
		sex:   "男",
		fight: 500000,
	}
	s := Person{"丁丁", 29, "女", 100000}
	fmt.Println(p)
	fmt.Printf("%+v\n", s)
}
