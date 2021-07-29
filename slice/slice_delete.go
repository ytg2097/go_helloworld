package main

import "fmt"

func main() {

	// 去头去尾的思想就是再切片
	deleteFront()
	deleteBack()
	deleteMid()
}

func deleteBack() {
	var s1 []int
	for i := 0; i < 10; i++ {
		s1 = append(s1, 2*i+1)
	}
	printS(s1)
	s1 = s1[:len(s1)-1]
	printS(s1)
}

func deleteFront() {
	var s1 []int
	for i := 0; i < 10; i++ {
		s1 = append(s1, 2*i+1)
	}
	printS(s1)
	s1 = s1[1:]
	printS(s1)

}

func deleteMid() {
	var s1 []int
	for i := 0; i < 10; i++ {
		s1 = append(s1, 2*i+1)
	}

	// 删除slice中的某个元素的思想是将slice当做一把尺子, 比如尺子从1到10, 删掉3就是从3位置折断
	// 然后将两节尺子叠到一起, 将4覆盖住原来3的位置
	// 1,2,3,4,5,6,7,8,9,10
	//	   3被盖住
	// 1,2,4,5,6,7,8,9,10
	// 换成代码就是s[:3] + s[4:]
	printS(s1)
	s1 = append(s1[:3], s1[4:]...)
	printS(s1)
}

func printS(s []int) {
	fmt.Printf("val = %v, len = %d, cap = %d \n", s, len(s), cap(s))
}
