package main

import "fmt"

func main() {
	var s1 []int
	//s1 = make([]int, 3) //构造切片
	s1 = append(s1, 1) //追加切片
	s1 = append(s1, 2)
	//s1 = append(s1, 3)
	//s1 = append(s1, 4)
	//s1 = append(s1, 5)
	printSlice(s1)

	s2 := make([]int, 3, 3)
	s2 = append(s2, 100) //自动扩充容量，倍数扩充*2
	printSlice(s2)

	copy(s2, s1)
	printSlice(s2)
	s2[0] = 10000
	printSlice(s2)
	printSlice(s1)
}

func printSlice(s []int) {
	fmt.Printf("len = %d, cap =%d, %v\n", len(s), cap(s), s)
}
