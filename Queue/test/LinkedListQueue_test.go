package test

import (
	"algorithm/Queue"
	"fmt"
	"testing"
)

func TestLinkedListQueue(t *testing.T) {
	LinkedListQueue := Queue.NewLinkedListQueue()
	for i := 0; i < 10; i++ {
		LinkedListQueue.Enqueue(i)
		fmt.Println(LinkedListQueue)
		if i%3 == 2 {
			LinkedListQueue.Dequeue()
			fmt.Println(LinkedListQueue)
		}
	}
}
