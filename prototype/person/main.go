package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address struct {
	Suite               int
	StreetAddress, City string
}

type Employee struct {
	Name   string
	Office *Address
}

/**
 * DeepCopy perform deep copy for prototyping purposes.
 * instead of loop and using copy() pointer we can use serialization like json or encode-decode like this
 */
func (e *Employee) DeepCopy() *Employee {
	b := bytes.Buffer{}
	encoded := gob.NewEncoder(&b)
	_ = encoded.Encode(e)

	decoded := gob.NewDecoder(&b)
	result := Employee{}
	_ = decoded.Decode(&result)

	return &result
}

/**
 * Predefined address object
 */
var mainOffice = Employee{
	"", &Address{0, "123 St", "Up"},
}

/**
 * Predefined address object
 */
var subOffice = Employee{
	"", &Address{0, "321 St", "Down"},
}

func newEmployee(proto *Employee, name string, suite int) *Employee {
	result := proto.DeepCopy()
	result.Name = name
	result.Office.Suite = suite
	return result
}

/**
 * NewMainOfficeEmployee is func prototype factory
 */
func NewMainOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&mainOffice, name, suite)
}

/**
 * NewSubOfficeEmployee is func prototype factory
 */
func NewSubOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&subOffice, name, suite)
}

func main() {
	jane := NewSubOfficeEmployee("Jane", 2)
	jhon := NewMainOfficeEmployee("Jhon", 1)

	fmt.Println(jane, jane.Office)
	fmt.Println(jhon, jhon.Office)
}
