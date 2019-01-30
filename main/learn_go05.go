package main

import "fmt"

func main() {
	a := 1
	//& 指针指向的内存地址
	//* 通过*间接访问目标对象
	var p *int = &a
	fmt.Println(*p)
	a++
	fmt.Println(*p)
	fmt.Println(1 << 10)

	if true {
		fmt.Println("if")
	}
	if g := 1; g > 0 {
		fmt.Println("if2")
	}

	//for语句的三种形式
	for i := 0; i < 3; i++ {

	}

	for a <= 3 {

	}

	for {
		break
	}

	//switch语句,不需要break
	switch a := 1; {
	case a >= 0:
		fallthrough
	case a > 12:

	default:
	}

	//跳转语句 goto break continue

LABEL1:
	for {
		for i := 0; i < 10; i++ {
			if i > 3 {
				break LABEL1
			}
		}
	}
	fmt.Println("break")

}
