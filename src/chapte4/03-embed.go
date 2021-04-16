//继承

package main

import "fmt"

type Person struct {
	name  string
	age   int
	sex   string
	fight int
}

type Superman struct {
	Person   //继承
	strength int
	speed    int
}

func (s *Superman) print() {
	fmt.Printf("name=%s,age=%d,sex=%s,fight=%d\n", s.name, s.age, s.sex, s.fight)
	fmt.Printf("stength=%d,speed=%d\n", s.strength, s.speed)
}
func (p *Person) setage(_age int) {
	p.age = _age
}
func main() {
	s1 := Superman{
		Person: Person{
			name:  "洒洒",
			age:   20,
			sex:   "男",
			fight: 500000,
		},
		strength: 100000,
		speed:    100,
	}
	s1.print()
	s1.setage(39)
	s1.print()
}
