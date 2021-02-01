package Queue

import "fmt"

type LinkedListQueue struct {
	head, tail *node
	size       int
}
type node struct {
	e    interface{}
	next *node
}

func NewNode(e interface{}, next *node) *node {
	return &node{
		e:    e,
		next: next,
	}
}

func NewNodeIsNil(e interface{}) *node {
	return &node{
		e:    e,
		next: nil,
	}
}
func NewLinkedListQueue() Queue {
	return &LinkedListQueue{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (l *LinkedListQueue) GetSize() int {
	return l.size
}

func (l *LinkedListQueue) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkedListQueue) Enqueue(e interface{}) {
	if l.tail == nil {
		l.tail = NewNodeIsNil(e)
		l.head = l.tail
	} else {
		l.tail.next = NewNodeIsNil(e)
		l.tail = l.tail.next
	}
	l.size++
}

func (l *LinkedListQueue) Dequeue() interface{} {
	if l.IsEmpty() {
		panic("Cannot dequeue from an empty queue")
	}
	retNode := l.head
	l.head = l.head.next
	retNode.next = nil
	if l.head == nil {
		l.tail = nil
	}
	l.size--
	return retNode.e
}

func (l *LinkedListQueue) GetFront() interface{} {
	if l.IsEmpty() {
		panic("Queue is  empty ")
	}
	return l.head.e
}

func (l *LinkedListQueue) String() string {
	var str string
	cur := l.head
	str += "Queue:front "
	//fmt.Println("kkkk :",cur.e)`
	for cur != nil {
		val := cur.e
		str += fmt.Sprintf("%#v", val) + "->"
		cur = cur.next
	}
	str += "nil tail"
	return str
}
