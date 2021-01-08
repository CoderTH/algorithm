package LinkedListStack

import (
	"algorithm/LinkedList"
	"fmt"
)

type ILinkedListStack interface {
	GetSize()int
	IsEmpty()bool
	Push(e interface{})
	Pop()interface{}
	Peek()interface{}
}

type LinkedListStack struct {
	linked LinkedList.ILinkedList
}

func NewLinkedListStarck()ILinkedListStack {
	return &LinkedListStack{
		linked: LinkedList.NewLinkedList(),
	}
}


func (l *LinkedListStack) GetSize() int {
	return l.linked.GetSize()
}

func (l *LinkedListStack) IsEmpty() bool {
	return l.linked.IsEmpty()
}
func (l *LinkedListStack) Push(e interface{}) {
	l.linked.AddFirst(e)
}
func (l *LinkedListStack) Pop() interface{} {
	return l.linked.RemoveFirst()
}
func (l *LinkedListStack) Peek() interface{} {
	return l.linked.GetFirst()
}

func (l *LinkedListStack) String() string {
	str := "Stack: top "+fmt.Sprint(l.linked)
	return str
}



