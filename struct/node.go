package main

import "fmt"

type TreeNode struct {
	value       int
	left, right *TreeNode
}

// 这个方法会被挂载到node'设个方法的接收者上
// 因为go是值传递, 所以如果改了值在外部是不生效的
// 如果想要实现setValue这类方法, 接受者可以定义为指针类型
func (node *TreeNode) setVal() {
	node.value = 1
}
func (node TreeNode) name() {
	fmt.Println(node.value)
}

func main() {

	var node TreeNode
	node.value = 1
	node.left = &TreeNode{value: 2}
	node.right = &TreeNode{value: 3}
	node.right.left = new(TreeNode)
	fmt.Println(node.right)

	// 数组
	nodes := []TreeNode{
		{},
		{},
		{},
	}
	fmt.Println(nodes)

}
