package main

import "fmt"

func main() {

	//	a:=1
	//	var p *int = &a
	//	fmt.Print(*p)

	//数组 array
	//var array1 [...]int
	array2 := [2]int{1: 1}
	var p *[2]int = &array2

	//for e := range array1 {
	//	e = 1;
	//	fmt.Println(e)
	//}
	fmt.Println(*p)

	// a := [2]int{1,2}

	// b := [1]int{1}

	//使用new是返回一个类型的指针
	p1 := new([10]int)
	fmt.Println(*p1)

	//多维数组
	a := [2][3]int{
		{1, 1, 1},
		{2, 2, 2},
	}
	fmt.Println(a)

	h := [...]int{5, 2, 6, 3, 9}
	length := len(h)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if h[i] < h[j] {
				temp := h[i]
				h[i] = h[j]
				h[j] = temp
			}
		}
	}
	fmt.Println(h)
}
