package main

import (
	"circularSingleLink/tools"
)

func main() {
	head := &tools.Node{} //初始化头结点
	hero1 := &tools.Node{
		Id:   1,
		Name: "穆",
	}
	hero2 := &tools.Node{
		Id:   2,
		Name: "阿鲁迪巴",
	}
	hero3 := &tools.Node{
		Id:   3,
		Name: "撒加",
	}
	tools.InsertNode(head, hero1)
	tools.InsertNode(head, hero2)
	tools.InsertNode(head, hero3)
	tools.ListNode(head)

	head = tools.DelNode(head, 1)
	tools.ListNode(head)
	head = tools.DelNode(head, 2)
	tools.ListNode(head)
	head = tools.DelNode(head, 2)
	tools.ListNode(head)
}
