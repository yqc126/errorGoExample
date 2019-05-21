package example

import (
	"fmt"
	"testing"
)

func TestSlice_1_1(t *testing.T) {
	//初始化时根据内容确定长度和容量
	slice := []int{1, 2, 3}
	fmt.Printf("slice len = %d ,cap = %d ,v = %v ,addr = %p \n", len(slice), cap(slice), slice, slice)
	//append 之后扩容导致底层的数组变化
	slice = append(slice, 4)
	fmt.Printf("slice len = %d ,cap = %d ,v = %v ,addr = %p \n", len(slice), cap(slice), slice, slice)
	//append 之后容量足够,不进行扩容,数组地址不变
	slice = append(slice, 5)
	fmt.Printf("slice len = %d ,cap = %d ,v = %v ,addr = %p \n", len(slice), cap(slice), slice, slice)
}

func TestSlice_1_2(t *testing.T) {
	//初始化时指定长度,不指定容量,容量跟长度
	slice := make([]int, 1)
	fmt.Printf("slice len = %d ,cap = %d ,v = %v ,addr = %p \n", len(slice), cap(slice), slice, slice)
	//append 之后扩容导致底层的数组变化
	slice = append(slice, 4)
	fmt.Printf("slice len = %d ,cap = %d ,v = %v ,addr = %p \n", len(slice), cap(slice), slice, slice)
	//append 之后容量足够,不进行扩容,数组地址不变
	slice = append(slice, 5)
	fmt.Printf("slice len = %d ,cap = %d ,v = %v ,addr = %p \n", len(slice), cap(slice), slice, slice)
}

func TestSlice_1_3(t *testing.T) {
	//初始化时指定长度和容量,此时的长度不能超过容量,否则会报错
	//这时切片的容量是指定的容量
	slice := make([]int, 2, 3)
	fmt.Printf("slice len = %d ,cap = %d ,v = %v ,addr = %p \n", len(slice), cap(slice), slice, slice)
	//append 之后容量足够,不进行扩容,数组地址不变
	slice = append(slice, 4)
	fmt.Printf("slice len = %d ,cap = %d ,v = %v ,addr = %p \n", len(slice), cap(slice), slice, slice)
	//append 之后扩容导致底层的数组变化
	slice = append(slice, 5)
	fmt.Printf("slice len = %d ,cap = %d ,v = %v ,addr = %p \n", len(slice), cap(slice), slice, slice)
}
