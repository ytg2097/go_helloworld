package main

import "fmt"

func main() {

	maxLen := longSubStrLen("askjskdfjenbrwbvsdv")
	fmt.Println(maxLen)
}

/**
寻找最长的不重复的子串  返回他的长度

思路:

滑动窗口
*/
func longSubStrLen(str string) int {

	// lastOccurred用于记录每个字符最后出现的位置
	lastOccurred := make(map[byte]int)
	start := 0
	maxLen := 0
	// 字符串转字节数组的方式   var arr []byte = []byte(str)
	for i, ch := range []byte(str) {

		// 如果ch在lastOccurred中出现过并且本次遍历中ch的位置不在窗口的左侧
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			// 没有在start之前出现过   那么说明发现了新的最大不重复子串
			// 移动窗口左侧到新发现的ch的右边
			start = lastOccurred[ch] + 1
		}
		// 如果start到当前遍历字符之间的长度大于之前的maxLen,  说明发现了更长的
		if i-start+1 > maxLen {
			maxLen = i - start + 1
		}

		// 不管怎样, lastOccurred中都要记录这个字符出现的位置
		lastOccurred[ch] = i
	}

	return maxLen
}
