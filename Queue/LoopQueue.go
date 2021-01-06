package Queue

import "fmt"

type ILoopQueue interface {
	GetCapacity()int
	IsEmpty()bool
	GetSize()int
	Enqueue(e interface{})
	Dequeue()interface{}
	GetFront()interface{}
}
//循环队列
type LoopQueue struct {
	data []interface{}
	front,tail int
	size int
}


func NewLoopQueue(capacity int) ILoopQueue {
	return &LoopQueue{
		data: make([]interface{},capacity+1),
	}
}
func (l *LoopQueue) GetCapacity() int {
	return len(l.data)-1
}

func (l *LoopQueue) IsEmpty() bool {
	return l.front == l.tail
}

func (l *LoopQueue) GetSize() int {
	return l.size
}
func (l *LoopQueue) Enqueue(e interface{}) {
	if (l.tail+1)%len(l.data)==l.front {
		l.resize(l.GetCapacity()*2)
	}
	l.data[l.tail] = e
	l.tail = (l.tail+1)%len(l.data)
	l.size++
}

func (l *LoopQueue) Dequeue() interface{} {
	if l.IsEmpty() {
		panic("Cannot dequeue from an empty queue.")
	}
	ret := l.data[l.front]
	l.front = (l.front+1)%len(l.data)
	l.size--
	if l.size == l.GetCapacity()/4 && l.GetCapacity()/2!=0{
		l.resize(l.GetCapacity()/2)
	}
	return ret
}

func (l *LoopQueue) resize(newCapacity int) {
	newData := make([]interface{},newCapacity+1)
	for i := 0; i < l.size; i++ {
		newData[i] = l.data[(l.front+i)%len(l.data)]
	}
	l.data = newData
	l.front = 0;
	l.tail = l.size
}
func (l *LoopQueue) GetFront() interface{} {
	if l.IsEmpty() {
		panic("Queue is empty")
	}
	return l.data[l.front]
}
func (l *LoopQueue)String()string{
	var str string
	arr := make([]interface{},0,10)
	for i := l.front; i != l.tail; i=(i+1)%len(l.data) {
		arr = append(arr, l.data[i])
	}
	str = fmt.Sprintf("Queue : size = %d,capacity = %d front %v tail \n",l.size,l.GetCapacity(),arr)
	return str
}





