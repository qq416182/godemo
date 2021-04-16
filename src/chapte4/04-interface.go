//多态
package main

import "fmt"

//定义动物接口
type Animal interface {
	sleeping()
	eating()
}

type Dog struct {
	color string
}
type Cat struct {
	color string
}

func (d Dog) sleeping() {
	fmt.Printf("%s's dog is sleeping\n", d.color)
}
func (d Dog) eating() {
	fmt.Printf("%s's dog is sleeping\n", d.color)
}

func (c Cat) sleeping() {
	fmt.Printf("%s's cat is sleeping\n", c.color)
}
func (c Cat) eating() {
	fmt.Printf("%s's cat is sleeping\n", c.color)
}

//工厂模式：流水线作业。可以生产猫，也可以生产狗
func factory(color string, animal string) Animal {
	switch animal {
	case "dog":
		return &Dog{color}
	case "cat":
		return &Cat{color}
	default:
		return nil //NULL
	}
}
func main() {
	c1 := Cat{"white"}
	c1.eating()
	d1 := Dog{"black"}
	d1.sleeping()
	fmt.Println("--------------")
	c2 := factory("yellow", "cat")
	c2.eating()
	d2 := factory("green", "dog")
	d2.sleeping()

}
