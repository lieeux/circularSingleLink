package tools

import "fmt"

type Node struct {
	Id   int
	Name string
	next *Node //指向后一个节点
}

func InsertNode(head *Node, newNode *Node) {
	if head.next == nil { //判断是不是第一个节点
		head.Id = newNode.Id
		head.Name = newNode.Name
		head.next = head //指向自己，形成环
	} else {
		temp := head //创建一个辅助节点
		for {        //找环尾
			if temp.next == head {
				break
			}
			temp = temp.next
		}
		temp.next = newNode //将链表末尾指向新节点
		newNode.next = head
	}
}

func DelNode(head *Node, id int) *Node { //删除节点
	if head.next == nil { //如果是空链
		fmt.Println("链表为空")
		return head
	}
	if head.next == head && head.Id == id { //如果只有一个节点且要删的就是这个节点
		head.next = nil
		return head
	}

	temp := head //创建一个辅助节点
	for {
		if temp.next == head && head.Id == id { //找了一圈发现要删的是头
			head = head.next
			temp.next = temp.next.next //把当前节点指向下个节点的下个节点
			break
		}
		if temp.next == head && head.Id != id { //找了一圈发现要删的不是头
			fmt.Println("Id不存在")
			break
		}
		if temp.next.Id == id { //这里不需要&& head.Id != id，因为如果是头，第一个判断就跳出来了；但如果把这个判断写在前面，就需要加这个附加条件
			temp.next = temp.next.next //把当前节点指向下个节点的下个节点
			break
		}
		temp = temp.next
	}
	return head //头结点改变后需要在主函数接收
}

func ListNode(head *Node) {
	temp := head
	if temp.next == nil {
		fmt.Println("链表为空")
		return
	}
	for {
		fmt.Printf("[%d %s]->", temp.Id, temp.Name) //输出当前节点
		if temp.next == head {                      //先判断
			break
		}
		temp = temp.next //再跳转
	}
	fmt.Println(head.Id)
}
