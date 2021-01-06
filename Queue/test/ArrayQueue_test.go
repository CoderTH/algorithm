package test

import (
	"algorithm/Queue"
	"fmt"
	"testing"
)

func TestArrayQueue(t *testing.T)  {
	ArrayQueue := Queue.NewArrayQueue(10)
	for i := 0; i < 10; i++ {
		ArrayQueue.Enqueue(i)
		fmt.Println(ArrayQueue)
		if i%3==2 { 
			ArrayQueue.Dequeue();
			fmt.Println(ArrayQueue)
		}
	}
}
