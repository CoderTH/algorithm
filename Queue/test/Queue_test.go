package test

import (
	"algorithm/Queue"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestQueue(t *testing.T) {
	opCount := 100000
	loopQueue := Queue.NewLoopQueue(10)
	queue := testQueue(loopQueue, opCount)
	fmt.Println("loopQueue time  : ", queue)
	ArrayQueue := Queue.NewArrayQueue(10)
	queue2 := testQueue(ArrayQueue, opCount)
	fmt.Println("ArrayQueue time  : ", queue2)
	LinkedListQueue := Queue.NewLinkedListQueue()
	queue3 := testQueue(LinkedListQueue, opCount)
	fmt.Println("LinkedListQueue time  : ", queue3)

}

func testQueue(queue Queue.Queue, opCount int) time.Duration {
	start := time.Now()
	for i := 0; i < opCount; i++ {
		queue.Enqueue(rand.Intn(100))
	}
	for i := 0; i < opCount; i++ {
		queue.Dequeue()
	}
	elapsed := time.Since(start)

	return elapsed
}
