package main

import "fmt"

func main() {
	//定义5个元素整形数组
	var five [5]int = [5]int{1, 2, 3, 4, 5}
	//定义3个元素的字符串数组
	//var three [3]string = [3]string{"one", "two", "three"}
	var three [3]string = [3]string{"one", "two"}

	fmt.Println(five)
	fmt.Println(three)
}
