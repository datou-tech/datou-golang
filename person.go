package main

import "fmt"

type personInterface interface {
	printName()
}

type person struct {
	name string
	age  int
}

func newPerson(n string, a int) person {
	p := person{name: n} // using named parameters allows you to construct without all the fields
	p.age = a
	return p
}

func (p person) printName() {
	fmt.Println(p.name)
}

func (p person) isAdult() bool {
	if p.age >= 18 {
		return true
	}
	return false
}
