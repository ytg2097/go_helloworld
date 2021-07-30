package main

import "fmt"

func main() {

	// go中map是hashmap
	// go中除了slice, map, func之外都可以作为map的key
	// go中自定义类型struct也可作为key, 只要field中不包含slice, map, func就可以
	m := map[int]string{
		1: "123",
		2: "345",
	}
	fmt.Println(m)

	m1 := map[int]map[string]string{
		1: {"age": "123", "name": "123"},
		2: {"age": "123", "name": "123"},
	}
	fmt.Println(m1)

	// 新建空map
	m2 := make(map[int]string)
	fmt.Println(m2)

	// 新建空map
	var m3 map[string]int
	fmt.Println(m3)

	// 遍历map
	for key, val := range m {
		fmt.Println(key, val)
	}

	// 即使key在map中不存在, 也不会返回null nil
	// 因为go的每一个变量都有一个Zero Value
	valOf3 := m[3]
	fmt.Println(valOf3)

	// 判断key是否在map中存在
	val, exist := m[3]
	fmt.Println(val, exist)

	// 删除key
	delete(m, 1)
	val1, exist1 := m[1]
	fmt.Println(val1, exist1)

	//如果要对map排序  需要借助slice, 将key都放到slice中, 然后对slice排序
}
