package main

import "fmt"

// 包内公共变量可以var ()定义  包内可见
var (
	name = "张三"
	age  = 18
)

func main() {

	// 函数内变量可以:= 赋值   类型自动推断
	gender := "男"
	fmt.Println(name, age, gender)
}
