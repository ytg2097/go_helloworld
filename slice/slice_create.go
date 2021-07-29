package main

import "fmt"

func main() {

	// 不从array创建, 直接定义一个slice  这里声明的类型是[]int不是[]int{}
	var s []int
	for i := 0; i < 100; i++ {
		// 当cap不足进行扩容时, cap是二倍扩容的
		s = append(s, 2*i+1)
		printSlice(s)
	}

	s1 := make([]int, 10)
	s2 := make([]int, 16)
	printSlice(s1)
	printSlice(s2)
}

func printSlice(s []int) {
	fmt.Printf("len = %d, cap = %d \n", len(s), cap(s))
}
