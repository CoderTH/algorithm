package Set

import "algorithm/LinkedList"

type LinkedListSet struct {
	list LinkedList.ILinkedList
}

func NewLinkedListSet() Set {
	return &LinkedListSet{
		list: LinkedList.NewLinkedList(),
	}
}

func (l *LinkedListSet) Add(e int) {
	if !l.list.Contains(e) {
		l.list.AddFirst(e)
	}
}

func (l *LinkedListSet) Remove(e int) {
	l.list.RemoveElement(e)
}

func (l *LinkedListSet) Contains(e int) bool {
	return l.list.Contains(e)
}

func (l *LinkedListSet) GetSize() int {
	return l.list.GetSize()
}

func (l *LinkedListSet) IsEmpty() bool {
	return l.list.IsEmpty()
}
