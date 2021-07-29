package main

import "fmt"

func main() {

	var s1 []int
	for i := 0; i < 10; i++ {
		s1 = append(s1, 2*i+1)
	}

	s2 := make([]int, 16)
	copy(s2, s1)
	printSliceV(s1)
	printSliceV(s2)
}

func printSliceV(s []int) {
	fmt.Printf("val = %v, len = %d, cap = %d \n", s, len(s), cap(s))
}
