package test

import (
	"algorithm/Queue"
	"fmt"
	"testing"
)

func TestLoopQueue(t *testing.T) {
	loopQueue := Queue.NewLoopQueue(10)
	for i := 0; i < 10; i++ {
		loopQueue.Enqueue(i)
		fmt.Println(loopQueue)
		if i%3==2 {
			loopQueue.Dequeue();
			fmt.Println(loopQueue)
		}
	}
}
