package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 30
	return &p
}

func main() {
	fmt.Println(newPerson(("Jon")))
	fmt.Println(person{"Prajwal", 23})
	s := person{"ABC", 22}
	sp := &s
	fmt.Println(sp.name)
	fmt.Println(sp.age)
}
