package main

import "fmt"

func main() {
	childMap := make(map[string]string)

	childMap["xiaoming"] = "男的"
	childMap["xiaohong"] = "女的"
	childMap["xiaojing"] = "女的"
	childMap["xiaodong"] = "男的"

	fmt.Println(childMap) //map[xiaodong:男的 xiaohong:女的 xiaojing:女的 xiaoming:男的]
}
