package example

import (
	"fmt"
	"testing"
)

type Animal interface {
	Name() string
	Age() int
}

type Cat struct {
	name string
	age  int
}

func (p *Cat) Name() string {
	return p.name
}

func (p *Cat) Age() int {
	return p.age
}

func TestInterface_1_1(t *testing.T) {
	var i interface{} = Cat{name: "tom", age: 100}
	p, ok := i.(Animal)
	if !ok {
		fmt.Println("not an animal")
		return
	}
	fmt.Println(p.Name(), p.Age())
}

func TestInterface_1_2(t *testing.T) {
	var i interface{} = &Cat{name: "tom", age: 100}
	p, ok := i.(*Animal)
	if !ok {
		fmt.Println("not an animal")
		return
	}
	fmt.Println((*p).Name(), (*p).Age())
}

var _ Animal = &Cat{}

func TestInterface_1_3(t *testing.T) {
	var i interface{} = &Cat{name: "tom", age: 100}
	p, ok := i.(Animal)
	if !ok {
		fmt.Println("not an animal")
		return
	}
	fmt.Println(Animal.Name(p), Animal.Age(p))
}

func TestInterface_1_4(t *testing.T) {
	var i interface{} = &Cat{name: "tom", age: 100}
	switch i.(type) {
	case Animal:
		fmt.Println("this is an animal")
	default:
		fmt.Printf("%T %v", i, i)
	}
}
