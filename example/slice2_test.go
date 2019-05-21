package example

import (
	"fmt"
	"testing"
)

func TestSlice_2_1(t *testing.T) {
	//初始化一个容量为5的切片
	baseSlice := []int{1, 2, 3, 4, 5}
	fmt.Printf("baseSlice len = %d ,cap = %d ,v = %v ,addr = %p \n", len(baseSlice), cap(baseSlice), baseSlice, baseSlice)
	//取其中的一段
	childSlice := baseSlice[1:2] //[2]
	fmt.Printf("childSlice len = %d ,cap = %d ,v = %v ,addr = %p \n", len(childSlice), cap(childSlice), childSlice, childSlice)
	//添加元素,因为容量足够没有进行扩容
	childSlice = append(childSlice, 6)
	//可以看到切片的地址跟baseSlice[1]的地址是一样的
	fmt.Println("baseSlice[1] addr = ", &baseSlice[1])
	fmt.Printf("childSlice len = %d ,cap = %d ,v = %v ,addr = %p \n", len(childSlice), cap(childSlice), childSlice, childSlice)
	//最后看到baseSlice变了,跟我们想的有点不一样
	fmt.Printf("baseSlice len = %d ,cap = %d ,v = %v ,addr = %p \n", len(baseSlice), cap(baseSlice), baseSlice, baseSlice)
}

func TestSlice_2_2(t *testing.T) {
	//初始化一个容量为5的切片
	baseSlice := []int{1, 2, 3, 4, 5}
	fmt.Printf("baseSlice len = %d ,cap = %d ,v = %v ,addr = %p \n", len(baseSlice), cap(baseSlice), baseSlice, baseSlice)
	//取其中的一段
	childSlice := baseSlice[1:2:2] //[2]
	fmt.Printf("childSlice len = %d ,cap = %d ,v = %v ,addr = %p \n", len(childSlice), cap(childSlice), childSlice, childSlice)
	//添加元素,因为容量足够没有进行扩容
	childSlice = append(childSlice, 6)
	//可以看到切片的地址跟baseSlice[1]的地址是一样的
	fmt.Println("baseSlice[1] addr = ", &baseSlice[1])
	fmt.Printf("childSlice len = %d ,cap = %d ,v = %v ,addr = %p \n", len(childSlice), cap(childSlice), childSlice, childSlice)
	//最后看到baseSlice没有发生变化
	fmt.Printf("baseSlice len = %d ,cap = %d ,v = %v ,addr = %p \n", len(baseSlice), cap(baseSlice), baseSlice, baseSlice)
}
