package main

import "fmt"

type myType string

func main() {
	value := myType("hello world")
	value.method()
	value.pointerMethod() //值类型自动转化为指针

	pointer := &value
	pointer.method() //指针类型自动转换为值类型
	pointer.pointerMethod()
}

func (m myType) method() {
	fmt.Println("method value")
}

func (m *myType) pointerMethod() {

	fmt.Println("pointer method value")
}
