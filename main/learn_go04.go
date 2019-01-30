package main

import "fmt"

//全局变量
var (
	named = "hello"
	sss   = "123"
)

//iota枚举计数器
const (
	Max_number = 100
	aa, bb     = 1, 2
	cc, dd
	monday = iota
)

//运算符
func cal() {
	fmt.Println(^2)
	fmt.Println(1 ^ 2)
	fmt.Println(!true)
	fmt.Println(1 << 10 << 10 >> 10 >> 10)

	//& | &^ ^
}

func main() {
	fmt.Println(len(sss))
	fmt.Println(monday)
	fmt.Println(Max_number)
	cal()
}
