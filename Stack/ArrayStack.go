package Stack

import (
	"algorithm/Array/Array"
	"fmt"
)

type arrayStack struct {
	array Array.IArray
}

//构造函数
func NewArrayStack(capacity int) Stack {
	return &arrayStack{
		array: Array.NewArray(10),
	}
}

func (a *arrayStack) GetSize() int {
	return a.array.GetSize()
}

func (a *arrayStack) IsEmpty() bool {
	return a.array.IsEmpty()
}

func (a *arrayStack) GetCapacity() int{
	return a.array.GetCapacity()
}

func (a *arrayStack) Push(e interface{}) {
	a.array.AddLast(e)
}

func (a *arrayStack) Pop() interface{} {
	return a.array.RemoveLast()
}

func (a *arrayStack) Peek() interface{} {
	return a.array.GetLast()
}

func (a *arrayStack)String()string{
	var str string
	str = fmt.Sprintf("Stack : %v top \n",a.array.GetArray())
	return str
}





