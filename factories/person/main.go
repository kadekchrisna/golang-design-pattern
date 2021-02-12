package main

import "fmt"

type person struct {
	name, email string
}

func (p *person) Hello() {
	fmt.Printf("I'm %s and hello\n", p.name)
}

// Person ...
type Person interface {
	Hello()
}

// NewPerson ...
func NewPerson(name string, email string) Person {
	return &person{name, email}
}

func main() {
	p := NewPerson("k", "k@gmail.com")
	p.Hello()
}
