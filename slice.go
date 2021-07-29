package main

import (
	"fmt"
)

// slice是一个视图对象   修改slice会同时修改掉原来的数组内容
func main() {

	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	// 切片为 start   end - 1
	// 切片的第二个数字不是取多少个   而是到第几个角标 - 1
	slice1 := arr[2:5]
	slice2 := arr[:5]
	slice3 := arr[2:]
	slice4 := arr[:]
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)

	// slice还可以在被切片   被切片之后底层的数组是相同的
	slice5 := slice4[2:]
	fmt.Println(slice5)

	s1 := arr[2:6]
	/*
		slice在对数组切片后  在对arr做视图的同时, 视图右边的没有纳进视图的部分也是slice可见的, 如果
		对s1再次slice时, 即使超出了s1的len, 但是只要没有超出底层数组的len, 那么依然可以slice成功

		slice的底层结构有三个属性
		ptr: 视图在原数组的index
		len: 视图的长度
		cap: arr.len - slice.ptr
	*/
	// 虽然slice操作可以成功获取  但是如果直接s1[5]是拿不到的  因为s1[5]是读操作   是找不到的
	s2 := s1[3:5]
	fmt.Println(s1)
	fmt.Println(s2)
}
