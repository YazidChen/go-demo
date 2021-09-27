package main

import "fmt"
import "reflect"

//接口
type Duck interface {
	Quack()
}

type Cat struct{}

func (c *Cat) Quack() {
	fmt.Println("meow")
}

//反射
func reflects() {
	a := "aaaaa"
	//接口值到反射对象转换
	fmt.Println("Type:", reflect.TypeOf(a))
	fmt.Println("Value:", reflect.ValueOf(a))
	//反射对象到接口追转换
	aa := reflect.ValueOf(a)
	b := aa.Interface().(string)
	fmt.Println(b)
}

func main() {
	//var c Duck = &Cat{}
	//c.Quack()
	reflects()
}
