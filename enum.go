package main

import "fmt"

func main() {

	// go的枚举类型就是一组常量
	const(
		cpp = 1
		java = 2
	// 如果一组常量中有一个没有显示赋值  那么它的值是上一个常量的值
		golang
		dd = 3

	)
	fmt.Println(cpp,java,golang, dd)

	// 自增值枚举iota 给常量赋值0   后面常量自增加一
	// iota 是不能给var使用的
	const(
		python = iota
		// 可以用 _ 跳过一个自增值
		_
		scala
		ee
		bb
	)
	fmt.Println(python,scala,ee,bb)

	// b kb mb gb tb pb

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	// iota还是 0开始自增   上面的计算过程等同于下面的
	fmt.Println( 1 << (10 * 0))
	fmt.Println( 1 << (10 * 1))
	fmt.Println( 1 << (10 * 2))
	fmt.Println( 1 << (10 * 3))
	fmt.Println(b,kb,mb,gb,tb,pb)
}
