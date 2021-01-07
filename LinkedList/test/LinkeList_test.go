package test

import (
	"algorithm/LinkedList"
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {
	linkedList := LinkedList.NewLinkedList()
	for i :=0;i<5;i++{
		linkedList.AddFirst(i)
		fmt.Println(linkedList)
	}
	linkedList.Add(2,666)
	fmt.Println(linkedList)
	linkedList.Remove(2)
	fmt.Println(linkedList)
	linkedList.RemoveFirst()
	fmt.Println(linkedList)
	linkedList.RemoveLast()
	fmt.Println(linkedList)
}
