package main

import "fmt"

type person struct {
	name string
	age  int
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Printf("My name is %s, and I am %d years old.\n", p.name, p.age)
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p := person{
		name: "Peter",
		age:  29,
	}
	saySomething(&p)
}
