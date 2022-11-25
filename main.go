package main

import (
	child1 "go-coding-pattrens.examples.com/Factory/Child1"
	child2 "go-coding-pattrens.examples.com/Factory/Child2"
	child3 "go-coding-pattrens.examples.com/Factory/Child3"
)

func main() {

	FactoryMethodDemo()

}

func FactoryMethodDemo() {
	// Run Time Polymorphisim
	c1 := child1.CreateObject()
	c2 := child2.CreateObject()
	c3 := child3.CreateObject()

	// c1 = c3

	c1.Print()
	c2.Print()
	c3.Print()

}
