package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	// if的条件中可以定义变量   作用域在if之内
	if file, err := ioutil.ReadFile("test.txt"); err != nil {
		fmt.Println(err)
	} else if len(err.Error()) == 1 {
		// printf方法可以格式化输出  配合%d %s等
		fmt.Printf("%s", file)
	} else {

	}

}
