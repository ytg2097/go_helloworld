package main

import "fmt"

// go里边的数组是值类型  不是引用类型!
// 所以要注意如果数组作为参数传递  那么如果传递之后数组中的值被修改   外部是不知道的   因为值传递是拷贝传递
// 所以如果传递数组  且要在修改数组后对外部可见的话可以传递数组的指针  也可以传递 slice  推荐使用slice
func main() {

	arr1 := [2]int{}
	arr2 := [3]int{1, 2, 3}
	// 不明确数量, 让编译器去数
	arr3 := [...]int{9, 9, 1, 3}
	grid := [...][3]int{{2, 4, 5}, {6, 7, 8}}
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(grid)
	loop(arr3)
}

func loop(arr [4]int) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// range返回两个结果  第一个是元素的角标
	// 第二个是当前元素
	for i, i2 := range arr {

		fmt.Println(i)
		fmt.Println(i2)
	}
}
