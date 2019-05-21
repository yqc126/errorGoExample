package example

import (
	"fmt"
	"testing"
)

func TestInterface_2_1(t *testing.T) {
	cat := GetCat()
	if cat == nil {
		fmt.Println("cat is nil")
	} else {
		fmt.Println(cat.Name(), cat.Age())
	}
}

func TestInterface_2_2(t *testing.T) {
	cat := GetCat()
	if cat == nil {
		fmt.Println("cat is nil")
	} else {
		fmt.Println(cat.Name(), cat.Age())
	}
	var i interface{} = cat
	if i == nil {
		fmt.Println("i is nil")
	} else {
		fmt.Println("i not nil")
	}
	fmt.Printf("%T %v\n", cat, cat)
}

func GetCat() *Cat {
	var cat *Cat
	//do something
	return cat
}

func GetObject() interface{} {
	var i int64 = 3
	return i
}

func TestInterface_2_3(t *testing.T) {
	age := GetObject()
	if age == nil {
		fmt.Println("age is nil")
	} else {
		fmt.Println("age not nil")
	}

	if age == 3 {
		fmt.Println("age is 3")
	} else {
		fmt.Println("age not 3")
	}
}
