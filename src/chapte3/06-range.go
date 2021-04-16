package main

import "fmt"

func main() {
	//var testMap map[string]string

	childMap := make(map[string]string)

	childMap["xiaoming"] = "男的"
	childMap["xiaohong"] = "女的"
	childMap["xiaojing"] = "女的"
	childMap["xiaodong"] = "男的"
	childMap["xiaofei"] = "男的"

	for k, v := range childMap {
		fmt.Println(k, " is ", v)
	}

	five := [5]int{1, 2, 3, 4, 5}
	for k, v := range five {
		fmt.Printf("a[%d] =%d\n", k, v)
	}
	//_表示占位
	for _, v := range five {
		fmt.Printf("%d\n", v)
	}

	ewei := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	for i, v := range ewei {
		//fmt.Printf("a[%d] =%d\n", i, v)
		for j, v1 := range v {
			fmt.Printf("a[%d,%d] =%d\n", i, j, v1)
		}
	}
}
