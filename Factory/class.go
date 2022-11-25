package factory

type IClass interface {
	Print()
}

type Class struct {
}

func (c *Class) CreateObject() IClass {
	return c
}

func (c *Class) Print() {

}
