package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	str := "hello杨同港!"
	//bytes := []byte(str)
	for i, by := range []byte(str) {
		fmt.Printf(" (%d, %X) ", i, by)
	}
	fmt.Println()

	// 中文是utf8编码的 解码后是三个字节, 而rune是unicode编码, 再编码又成四个字节
	for i, ru := range []rune(str) {
		fmt.Printf(" (%d, %X) ", i, ru)
	}
	fmt.Println()

	// 中文是utf8编码的 解码后是三个字节, 而rune是unicode编码, 再编码又成四个字节
	for i, ru := range str {
		fmt.Printf(" (%d, %X) ", i, ru)
	}

	fmt.Println(len(str))
	fmt.Println(len([]rune(str)))
	fmt.Println(utf8.RuneCountInString(str))

	bytes := []byte(str)
	for len(bytes) > 0 {
		// 不管bytes有多长, 这个方法一次只解码一个rune,  返回解码后的rune, 以及这个rune的字节数
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c", ch)
	}

	fmt.Println()

	for i, r := range []rune(str) {
		fmt.Printf("%d,%c", i, r)
	}
}
