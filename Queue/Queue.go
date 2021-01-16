package Queue

import (
	"algorithm/Array/Array"
	"fmt"
)

type Queue interface {
	GetSize() int
	IsEmpty()bool
	Enqueue(e interface{})
	Dequeue()interface{}
	GetFront()interface{}
}

type ArrayQueue struct {
    array Array.IArray
}

func NewArrayQueue(capacity int) Queue {
	return &ArrayQueue{
		array: Array.NewArray(10),
	}
}
func (a *ArrayQueue) GetSize() int {
	return a.array.GetSize()
}

func (a *ArrayQueue) IsEmpty() bool {
	return a.array.IsEmpty()
}

func (a *ArrayQueue) Enqueue(e interface{}) {
	a.array.AddLast(e)
}

func (a *ArrayQueue) Dequeue() interface{} {
	return a.array.RemoveFirst()
}

func (a *ArrayQueue) GetFront() interface{} {
	return a.array.GetFirst()
}

func (a *ArrayQueue)String()string{
	var str string
	str = fmt.Sprintf("Queue : front %v tail \n",a.array.GetArray())
	return str
}