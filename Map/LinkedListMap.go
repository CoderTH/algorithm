package Map

import "fmt"

type LinkedListMap struct {
	dummyHead *node
	size      int
}

type node struct {
	k, v interface{}
	next *node
}

func NewNode(k, v interface{}, next *node) *node {
	return &node{
		k:    k,
		v:    v,
		next: next,
	}
}

func NewLinkedListMap() Map {
	return &LinkedListMap{
		dummyHead: NewNode(nil, nil, nil),
	}
}

func (l *LinkedListMap) Add(key interface{}, value interface{}) {
	node := l.getNode(key)
	if node == nil {
		l.dummyHead.next = NewNode(key, value, l.dummyHead.next)
		l.size++
	} else {
		node.v = value
	}
}

func (l *LinkedListMap) Remove(key interface{}) interface{} {
	prev := l.dummyHead
	for prev.next != nil {
		if prev.next.k == key {
			break
		}
		prev = prev.next
	}
	if prev.next != nil {
		delNode := prev.next
		prev.next = delNode.next
		delNode.next = nil
		return delNode.v
	}
	return nil
}

func (l *LinkedListMap) Contains(key interface{}) bool {
	return l.getNode(key) != nil
}

func (l *LinkedListMap) Get(Key interface{}) interface{} {
	node := l.getNode(Key)
	if node == nil {
		return nil
	} else {
		return node.v
	}
}

func (l *LinkedListMap) Set(key interface{}, value interface{}) {
	node := l.getNode(key)
	if node != nil {
		node.v = value
	} else {
		fmt.Println(key, "doesn't exist!")
	}
}

func (l *LinkedListMap) GetSize() int {
	return l.size
}

func (l *LinkedListMap) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkedListMap) getNode(key interface{}) *node {
	cur := l.dummyHead.next
	for cur != nil {
		if cur.k == key {
			return cur
		}
		cur = cur.next
	}
	return nil
}
