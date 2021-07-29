package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("test.txt");

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file);

	// go的for可以没有起始语句和递增语句   与while一样只有一个结束条件就可以
	// 甚至结束条件也没有   直接  for {}  死循环
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
