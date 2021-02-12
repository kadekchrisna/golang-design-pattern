package main

import (
	"fmt"
)

// Person ...
type Person struct {
	City    string
	Street  string
	ZipCode string

	Job     string
	Company string
	Salary  float64
}

// PersonBuilder ...
type PersonBuilder struct {
	Person *Person
}

// NewPersonBuilder ...
func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&Person{}}
}

// Build ...
func (pb *PersonBuilder) Build() *Person {
	return pb.Person
}

// Works ...
func (pb *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*pb}
}

// Lives ...
func (pb *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*pb}
}

// PersonAddressBuilder ...
type PersonAddressBuilder struct {
	PersonBuilder
}

// At ...
func (it *PersonAddressBuilder) At(street string) *PersonAddressBuilder {
	it.Person.Street = street
	return it
}

// In ...
func (it *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	it.Person.City = city
	return it
}

// ZipCode ...
func (it *PersonAddressBuilder) ZipCode(zip string) *PersonAddressBuilder {
	it.Person.ZipCode = zip
	return it
}

// PersonJobBuilder ...
type PersonJobBuilder struct {
	PersonBuilder
}

// As ...
func (it *PersonJobBuilder) As(as string) *PersonJobBuilder {
	it.Person.Job = as
	return it
}

// At ...
func (it *PersonJobBuilder) At(company string) *PersonJobBuilder {
	it.Person.Company = company
	return it
}

// Earning ...
func (it *PersonJobBuilder) Earning(salary float64) *PersonJobBuilder {
	it.Person.Salary = salary
	return it
}

func main() {
	p := NewPersonBuilder()
	p.Lives().
		At("sesame street").
		In("cloud").
		ZipCode("123123").
		Works().
		As("BE").
		At("OWN CORP").
		Earning(2000)
	person := p.Build()
	fmt.Printf("%+v", *person)
}
