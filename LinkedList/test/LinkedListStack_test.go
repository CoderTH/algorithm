package test

import (
	"algorithm/LinkedList/LinkedListStack"
	"fmt"
	"testing"
)

func TestLinkedListStack(t *testing.T) {
	linkedListStack := LinkedListStack.NewLinkedListStarck()
	for i := 0; i < 5; i++ {
		linkedListStack.Push(i)
		fmt.Println(linkedListStack)
	}
	linkedListStack.Pop()
	fmt.Println(linkedListStack)
}
