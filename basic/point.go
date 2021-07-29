package main

import "fmt"

func main() {

	// 交换两个数字的值
	a, b := 3, 4
	// 一个值得引用用 & 获取   比如   pa *int = &a  意思是  pa是指针类型  它是a的指针
	swap(&a, &b)
	fmt.Println(a, b)

	operationPointer()
}

func operationPointer() {

	var a uint = 1
	var pa *uint = &a
	fmt.Println(a, pa)

}

// 指针类型都是*开头
func swap(a, b *int) {
	*b, *a = *a, *b
}
