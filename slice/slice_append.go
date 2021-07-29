package main

import "fmt"

func main() {

	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]       // s1 = 	   2, 3, 4, 5 _6,_7
	s2 := s1[3:5]        // s2 = 		        5, 6,_7
	s3 := append(s2, 10) // s3 = 			    5, 6, 10
	s4 := append(s3, 11) // s4 = 			 	5, 6, 10, 11
	s5 := append(s4, 12) // s5 = 			    5, 6, 10, 11, 12
	fmt.Println(s3, s4, s5)
	fmt.Println(arr)
	// append不会影响原slice, 他会返回一个新的slice:  slice长度 + 1 , 尾部值为新添加的元素
	// append的元素是在slice的末尾添加的, 只要添加的元素位置没有超过slice的cap, 那么这个元素就会直接映射到原数组所在的位置上
	// 如果append的元素超过了slice的cap, 那么go就会从原来的arr复制一个出来, 长度更长, 然后把新元素放到上面去
	// 当arr改变之后, slice的ptr,cap, len都会改变, 所以会返回一个新的slice
	//   原来的arr如果没有使用的话会被垃圾回收
}
