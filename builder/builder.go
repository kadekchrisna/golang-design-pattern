package main

import "fmt"

/**
 * Builder pattern used to build complex object by chaining method
 * channing mehtod also called builer fluent which return the reciever in order to allowing chaining
 * it is recomended for not-so-many-fields object
 *
 * if there is object that so complex it need more than one builder object that
 * can work together to build it
 *
 * i.e.
 * Like Person object it can be seperated by where he Lives and what he does for living and etc.
 */
func main() {
	fmt.Println("BUILDER")
}
