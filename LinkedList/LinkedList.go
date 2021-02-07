package LinkedList

import "fmt"

type ILinkedList interface {
	GetSize() int
	IsEmpty() bool
	AddFirst(e interface{})
	Add(index int, e interface{})
	AddLast(e interface{})
	Get(index int) interface{}
	GetFirst() interface{}
	GetLast() interface{}
	Set(index int, e interface{})
	Contains(e interface{}) bool
	Remove(index int) interface{}
	RemoveFirst() interface{}
	RemoveLast() interface{}
	RemoveElement(e interface{})
}
type LinkedList struct {
	dummyHead *node
	size      int
}

type node struct {
	e    interface{}
	next *node
}

//node的构造函数
func NewNodeIsNil(e interface{}) *node {
	return &node{
		e:    e,
		next: nil,
	}
}
func NewNode(e interface{}, next *node) *node {
	return &node{
		e:    e,
		next: next,
	}
}
func NewLinkedList() ILinkedList {
	return &LinkedList{
		dummyHead: NewNode(nil, nil),
		size:      0,
	}
}

//链表是否为空
func (l *LinkedList) IsEmpty() bool {
	return l.size == 0
}

//链表元素的个数
func (l *LinkedList) GetSize() int {
	return l.size
}

//在指定index位置插入节点
func (l *LinkedList) Add(index int, e interface{}) {
	if index < 0 || index > l.size {
		panic("Add fail. Illegal index.")
	}
	prev := l.dummyHead
	for i := 0; i < index; i++ {
		prev = prev.next
	}
	//node := NewNode(e,nil)
	//node.next = prev.next
	//prev.next = node
	prev.next = NewNode(e, prev.next)
	l.size++
}

//在链表的头部添加元素e
func (l *LinkedList) AddFirst(e interface{}) {
	l.Add(0, e)
}

//在最后添加
func (l *LinkedList) AddLast(e interface{}) {
	l.Add(l.size, e)
}

func (l *LinkedList) Get(index int) interface{} {
	if index < 0 || index > l.size {
		panic("Add fail. Illegal index.")
	}
	cur := l.dummyHead.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	return cur.e
}

//获取链表第一个元素
func (l *LinkedList) GetFirst() interface{} {
	return l.Get(0)
}

//获取链表最后一个元素
func (l *LinkedList) GetLast() interface{} {
	return l.Get(l.size)
}

//修改index位置的数据
func (l *LinkedList) Set(index int, e interface{}) {
	if index < 0 || index >= l.size {
		panic("Set faild. Illegal index.")
	}
	cur := l.dummyHead.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	cur.e = e
}

//查找在链表中是否存在e这个元素
func (l *LinkedList) Contains(e interface{}) bool {
	cur := l.dummyHead.next
	for cur.next != nil {
		if cur.e == e {
			return true
		}
		cur = cur.next
	}
	return false

}

//删除指定节点并返回删除的数据
func (l *LinkedList) Remove(index int) interface{} {
	if index < 0 || index > l.size {
		panic("remove failed.Index is illegal.")
	}
	prev := l.dummyHead
	for i := 0; i < index; i++ {
		prev = prev.next
	}
	deleteNode := prev.next
	prev.next = deleteNode.next
	deleteNode.next = nil
	l.size--
	return deleteNode.e
}

//删除链表中第一个元素
func (l *LinkedList) RemoveFirst() interface{} {
	return l.Remove(0)
}

//删除链表中最后一个元素
func (l *LinkedList) RemoveLast() interface{} {
	return l.Remove(l.size - 1)
}
func (l *LinkedList) String() string {
	var str string
	cur := l.dummyHead.next
	//fmt.Println("kkkk :",cur.e)`
	for cur != nil {
		val := cur.e
		str += fmt.Sprintf("%#v", val) + "->"
		cur = cur.next
	}
	str += "nil"
	return str
}

func (l *LinkedList) RemoveElement(e interface{}) {
	prev := l.dummyHead
	for prev.next != nil {
		if prev.next.e == e {
			break
		}
		prev = prev.next
	}
	if prev.next != nil {
		delNode := prev.next
		prev.next = delNode.next
		delNode.next = nil
	}
}
