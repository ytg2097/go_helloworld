package main

import "fmt"

func main() {

	q, r := unnamedDiv(20, 20)
	div, i := namedDiv(2, 33)

	// 如果函数返回了两个返回值  只想用其中一个的话可以把另一个用 _ 接收
	di, _ := namedDiv(2, 33)

	fmt.Println(q, r)
	fmt.Println(div, i)
	fmt.Println(di)

	if num, err := returnError(1, 2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(num)
	}

	apply(mapper, "我好")
}

// go的函数可以返回两个返回值
func unnamedDiv(a, b int) (int, int) {

	return a / b, a % b
}

// 这种方式尽量不用  可读性不好
func namedDiv(a, b int) (q, r int) {

	q = a / b
	r = a % b
	return
}

// error也是类型的一种   相当于声明了一个  throws Exception
func returnError(a, b int) (int, error) {

	// 返回一个error不会中断程序运行
	// 而panic则会中断程序的运行
	return 0, fmt.Errorf("error info %s", "错误")
}

// 函数的参数可以接受一个func   func可以是lamdba也是命名的func   是可以通过反射拿到func信息的
// go中函数是第一公民
func apply(mapper func(a string) string, str string) string {
	return mapper(str)
}

func mapper(str string) string {
	fmt.Println("丢弃掉原字符串" + str)
	fmt.Println(" 返回新字符串    你好")
	return "你好"
}
