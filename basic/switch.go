package main

func main() {

	// println是语言内置方法  他会输出红色 也就是error
	// fmt.Println是fmt包中的方法   会输出到标准
	println(eval("b"))

}

func switchNoCondition(ab string) int {

	// switch可以不写表达式
	// 在case中写表达式
	// 可以理解为换了语法的 if else
	switch {
	case ab == "c":
		return 1
	}
	return 2
}

func eval(abc string) int {

	var (
		result int
	)

	// go会自动break   如果不想break可以falltgrough
	switch abc {
	case "a":
		result = 1
	default:
		result = 1000
		panic("ssss")
	}

	return result
}
