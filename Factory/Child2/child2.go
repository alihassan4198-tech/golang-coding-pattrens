package child2

import (
	"fmt"

	factory "go-coding-pattrens.examples.com/Factory"
)

type child2St struct {
	Name string
}

func CreateObject() factory.IClass {
	c := child2St{}
	c.Name = "child 2"
	return &c
}

func (c2 *child2St) Print() {
	fmt.Println("printing child name : " + c2.Name)
}
