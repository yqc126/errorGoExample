package example

import (
	"fmt"
	"testing"
	"time"
)

func TestGoroutine_1_1(t *testing.T) {

	//错误的处理方式
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover")
		}
	}()
	go func() {
		//这里发生panic
		Say()
	}()

	time.Sleep(time.Second)
	fmt.Println("i am alive")
}

func Say() {
	fmt.Println("say somethine")
	panic("panic")
}

func TestGoroutine_1_2(t *testing.T) {
	go func() {
		//处理协程里面的panic
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("recover")
			}
		}()
		//这里发生panic
		Say()
	}()

	time.Sleep(time.Second)
	fmt.Println("i am alive")
}
