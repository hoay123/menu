package main

import "fmt"

type ListNode struct {
    Val interface{}
    next *ListNode
    lst *SingleList
}


//遍历
func (node *ListNode)Next() *ListNode {
	return node.next
}

type SingleList struct {
	head *ListNode
	len int
}

//单链表初始化
func (lst *SingleList)Init(){
	lst.head = new(ListNode)
	lst.head.lst = lst
	lst.len = 0
}

//创建单链表
func New() *SingleList {
	lst := new(SingleList)
	lst.Init()
	return lst
}

//清空链表
func (lst *SingleList)Clear(){
	lst.Init()
}

//返回第一个结点
func (lst *SingleList)Front() *ListNode {
	return lst.head.Next()
}

//返回尾结点
func (lst *SingleList)Back() *ListNode {
	if lst.len == 0{
		return nil
	}
	cur := lst.head.Next()
	for{
		if cur.Next() == nil{
			return cur
		}
		cur = cur.Next()
	}
}

//插入
func (lst *SingleList)insert(e, at *ListNode) *ListNode{
	lst.len++
	e.next = at.next
	at.next = e
	e.lst = lst
	return e
}

//指定元素后插
func (lst *SingleList)InsertAfter(val interface{}, mark *ListNode) *ListNode {
	if mark == nil{
		return nil
	}
	if mark.lst == lst{
		return lst.insert(&ListNode{Val:val},mark)
	}
	return nil
}

//尾插
func (lst *SingleList)PushBack(val interface{}) *ListNode{
	end := lst.Back()
	if end == nil{
		return lst.InsertAfter(val,lst.head)
	}else {
		return lst.InsertAfter(val,end)
	}
}

//头插
func (lst *SingleList)PushFront(val interface{}) *ListNode {
	return lst.InsertAfter(val,lst.head)
}

//删除
func (lst *SingleList)remove(e *ListNode) *ListNode{
	nextOne := e.next
	if nextOne == nil{
		return nil
	}
	lst.len--
	e.next = e.next.next
	nextOne.lst = nil
	return nextOne
}

//遍历
func (lst *SingleList)ShowList()  {
	for cur := lst.Front();cur!=nil;cur = cur.Next(){
		fmt.Print(cur.Val," ")
		cur = cur.Next()
	}
	fmt.Println()
}

//单链表长度
func (lst *SingleList)Len() int {
	return lst.len
}
