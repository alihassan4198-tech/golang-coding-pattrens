package child1

import (
	"fmt"

	factory "go-coding-pattrens.examples.com/Factory"
)

type child1St struct {
	Name string
}

func CreateObject() factory.IClass {
	c := child1St{}
	c.Name = "child 1"
	return &c
}

func (c1 *child1St) Print() {
	fmt.Println("printing child name : " + c1.Name)
}
