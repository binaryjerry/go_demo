package main

import (
	std "fmt"
	"go/types"
)

const Democonst = "helloworld"

const (
	PI = 3.14
)

//全局变量
var (
	name = ""
)

//一般类型
type ()

func main() {
	std.Printf(Democonst)
}

//Go的基本类型
type flag bool

//无符号
type number uint

//有符号
type number2 int64

//uint8别名
type byte1 byte

//int32别名
type r rune

//浮点型
type a float32

type b float64
type c complex64
type d complex128

//引用类型
type f types.Slice
