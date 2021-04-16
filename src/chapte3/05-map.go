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

	fmt.Println(childMap) //map[xiaodong:男的 xiaohong:女的 xiaojing:女的 xiaoming:男的]
	fmt.Println(childMap["xiaojing"])

	val := childMap["xiaojing"]
	fmt.Println("xiaojing is ", val)
	val = childMap["xiaofei"]
	fmt.Println("xiaofei is ", val)
	val, ok := childMap["xiaofei"]
	if ok {
		fmt.Println("xiaofei is ", val)
	} else {
		fmt.Println("xiaofei is not exists.")
	}
}
