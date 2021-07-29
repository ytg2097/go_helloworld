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

	/*
		slice在对数组切片后  在对arr做视图的同时, 视图右边的没有纳进视图的部分也是slice可见的, 如果
		对s1再次slice时, 即使超出了s1的len, 但是只要没有超出底层数组的len, 那么依然可以slice成功

		slice的底层结构有三个属性
		ptr: 视图在原数组的index
		len: 视图的长度
	*/
	// 可以这么理解  slice[x:y]会拥有数组的从x开始直到结束的所有元素,   y是不能超过数组的长度的
	// 				但是切片之后, 我们只可见数组的x:y部分的元素
	//				这个x在slice中就是ptr, x到数组尾部的长度成为capacity, slice的长度len就是y - x
	//				如果再次对slice进行切片, x是相对于slice的,而y还是相对于原数组的
	// 明白以上几点之后再去看slice的操作会清晰很多

	// arr = 				0,1,2,3,4,5,6,7
	s1 := arr[2:6] // s1 = 	2,3,4,5
	s2 := s1[3:9]  //  s2 =       	  5,6
	fmt.Println(s1)
	fmt.Println(s2)

}
