package example

import (
	"fmt"
	"testing"
)

//定义一个修改slice的函数
func ModifySlice(slice []int) {
	fmt.Printf("in ModifySlice len = %d ,cap = %d ,v = %v ,addr = %p \n", len(slice), cap(slice), slice, slice)
	slice = append(slice, 7, 8, 9)
	fmt.Printf("in ModifySlice len = %d ,cap = %d ,v = %v ,addr = %p \n", len(slice), cap(slice), slice, slice)
	slice[0] = 123
	fmt.Printf("in ModifySlice len = %d ,cap = %d ,v = %v ,addr = %p \n", len(slice), cap(slice), slice, slice)
}

func TestSlice_3_1(t *testing.T) {
	//定义一个slice
	slice := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("out ModifySlice len = %d ,cap = %d ,v = %v ,addr = %p \n", len(slice), cap(slice), slice, slice)
	ModifySlice(slice)
	//发现7,8,9并没有被append进来,而且slice[0]也没有变化
	fmt.Printf("out ModifySlice len = %d ,cap = %d ,v = %v ,addr = %p \n", len(slice), cap(slice), slice, slice)
}

func TestSlice_3_2(t *testing.T) {
	//初始化一个长度为0,容量为5的slice
	slice := make([]int, 0, 5)
	slice = append(slice, 1, 2)
	fmt.Printf("out ModifySlice len = %d ,cap = %d ,v = %v ,addr = %p \n", len(slice), cap(slice), slice, slice)
	ModifySlice(slice)
	//发现7,8,9并没有被append进来,而且slice[0]变化了
	fmt.Printf("out ModifySlice len = %d ,cap = %d ,v = %v ,addr = %p \n", len(slice), cap(slice), slice, slice)
}

func TestSlice_3_3(t *testing.T) {
	//初始化一个长度为0,容量为5的slice
	slice := make([]int, 0, 5)
	slice = append(slice, 1, 2, 3, 4)
	fmt.Printf("out ModifySlice len = %d ,cap = %d ,v = %v ,addr = %p \n", len(slice), cap(slice), slice, slice)
	ModifySlice(slice)
	//发现7,8,9并没有被append进来,而且slice[0]没有变化
	fmt.Printf("out ModifySlice len = %d ,cap = %d ,v = %v ,addr = %p \n", len(slice), cap(slice), slice, slice)
}

type SliceTest struct {
	Name  string
	Slice []int
}

func CopySliceTest(test SliceTest) SliceTest {
	test.Slice = append(test.Slice, 8, 9, 10)
	test.Slice[0] = 123
	return test
}

func TestSlice_3_4(t *testing.T) {
	oldStruct := SliceTest{Name: "test", Slice: []int{1, 2, 3, 4}}
	fmt.Printf("oldStruct = %v ,addr = %p\n", oldStruct, &oldStruct)
	newStruct := CopySliceTest(oldStruct)
	//Slice变化了
	fmt.Printf("oldStruct = %v ,addr = %p\n", oldStruct, &oldStruct)
	//newStruct的地址跟oldStruct不一样,但是Slice的数据一样
	fmt.Printf("newStruct = %v ,addr = %p\n", newStruct, &newStruct)
}

func TestSlice_3_5(t *testing.T) {
	slice := make([]int, 4, 5)
	oldStruct := SliceTest{Name: "test", Slice: slice}
	fmt.Printf("oldStruct = %v ,addr = %p ,slice = %v ,slice addr = %p \n", oldStruct, &oldStruct, oldStruct.Slice, oldStruct.Slice)
	newStruct := CopySliceTest(oldStruct)
	fmt.Printf("oldStruct = %v ,addr = %p ,slice = %v ,slice addr = %p \n", oldStruct, &oldStruct, oldStruct.Slice, oldStruct.Slice)
	fmt.Printf("newStruct = %v ,addr = %p ,slice = %v ,slice addr = %p \n", newStruct, &newStruct, newStruct.Slice, newStruct.Slice)
	fmt.Printf("slice =%v ,addr = %p\n", slice, slice)
}

func TestSlice_3_6(t *testing.T) {
	slice := make([]int, 1, 5)
	oldStruct := SliceTest{Name: "test", Slice: slice}
	fmt.Printf("oldStruct = %v ,addr = %p ,slice = %v ,slice addr = %p \n", oldStruct, &oldStruct, oldStruct.Slice, oldStruct.Slice)
	newStruct := CopySliceTest(oldStruct)
	fmt.Printf("oldStruct = %v ,addr = %p ,slice = %v ,slice addr = %p \n", oldStruct, &oldStruct, oldStruct.Slice, oldStruct.Slice)
	fmt.Printf("newStruct = %v ,addr = %p ,slice = %v ,slice addr = %p \n", newStruct, &newStruct, newStruct.Slice, newStruct.Slice)
	fmt.Printf("slice =%v ,addr = %p\n", slice, slice)
}
