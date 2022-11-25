package child3

import (
	"fmt"

	factory "go-coding-pattrens.examples.com/Factory"
)

type child3St struct {
	Name string
}

func CreateObject() factory.IClass {
	c := child3St{}
	c.Name = "child 3"
	return &c
}

func (c3 *child3St) Print() {
	fmt.Println("printing child name : " + c3.Name)
}
