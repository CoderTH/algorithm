package Array

import "fmt"

type IArray interface {
	GetSize() int
	GetCapacity() int
	IsEmpty() bool
	AddLast(e interface{})
	Add(index int, e interface{})
	AddFirst(e interface{})
	String() string
	Get(index int) interface{}
	Set(index int, e interface{})
	Contains(e interface{}) bool
	Find(e interface{}) int
	Remove(index int) interface{}
	RemoveFirst() interface{}
	RemoveLast() interface{}
	RemoveElement(e interface{}) bool
	GetLast() interface{}
	GetFirst() interface{}
	GetArray() []interface{}
	Swap(i, j int)
}

//不支持自定义类型
type Array struct {
	data []interface{}
	size int
}

//构造函数，传入数组的容量cap构造Array
func NewArray(cap int) IArray {
	return &Array{
		data: make([]interface{}, cap),
		size: 0,
	}
}

func (a *Array) Swap(i, j int) {
	if i < 0 || i >= a.size || j < 0 || j >= a.size {
		panic("Index is illegal.")
	}
	a.data[i], a.data[j] = a.data[j], a.data[i]
}

func (a *Array) GetArray() []interface{} {
	return a.data[:a.size]
}

func (a *Array) GetLast() interface{} {
	return a.Get(a.size - 1)
}

func (a *Array) GetFirst() interface{} {
	return a.Get(0)
}

//获取数组元素个数
func (a *Array) GetSize() int {
	return a.size
}

//获取数组的容量
func (a *Array) GetCapacity() int {
	return len(a.data)
}

//判断动态数组是否为空
func (a *Array) IsEmpty() bool {
	return a.size == 0
}

//添加新的元素
func (a *Array) AddLast(e interface{}) {
	a.Add(a.size, e)
}

//添加到头
func (a *Array) AddFirst(e interface{}) {
	a.Add(0, e)
}

//在第index个位置上插入一个新的元素
func (a *Array) Add(index int, e interface{}) {

	if index < 0 || index > a.size {
		panic("AddLst failed.Required index>=0and")
	}
	if a.size == len(a.data) {
		a.resize(2 * a.size)
	}
	for i := a.size - 1; i >= index; i-- {
		a.data[i+1] = a.data[i]
	}
	a.data[index] = e
	a.size++
}

func (a *Array) String() string {
	var str string
	str = fmt.Sprintf("Array : size = %d,capacity = %d\n", a.size, len(a.data))
	str = fmt.Sprintf("%s%v\n", str, a.data[0:a.size])
	return str
}

//获取索引位置的值
func (a *Array) Get(index int) interface{} {
	if index < 0 || index >= a.size {
		panic("Get failed. Index is illegal")
	}
	return a.data[index]
}

//设置指定索引位置的值
func (a *Array) Set(index int, e interface{}) {
	if index < 0 || index >= a.size {
		panic("Set failed. Index is illegal")
	}
	a.data[index] = e
}

//判断是否有e这个元素
func (a *Array) Contains(e interface{}) bool {
	for i := 0; i < a.size; i++ {
		if a.data[i] == e {
			return true
		}
	}
	return false
}

//查找数组中元素e所在的索引，如果不存在元素e，则返回-1
func (a *Array) Find(e interface{}) int {
	for i := 0; i < a.size; i++ {
		if a.data[i] == e {
			return i
		}
	}
	return -1
}

//从数组中删除index位置的元素，返回删除的元素
func (a *Array) Remove(index int) interface{} {
	if index < 0 || index > a.size {
		panic("Remove failed.Required index>=0and")
	}
	res := a.data[index]
	for i := index; i < a.size-1; i++ {
		a.data[i] = a.data[i+1]
	}
	a.size--
	if a.size == len(a.data)/4 && len(a.data)/2 != 0 {
		a.resize(len(a.data) / 2)
	}
	return res
}

//删除数组中第一个元素，返回删除的元素
func (a *Array) RemoveFirst() interface{} {
	return a.Remove(0)
}

//删除数组中最后一个元素，返回删除元素
func (a *Array) RemoveLast() interface{} {
	return a.Remove(a.size - 1)
}

//从数组数组中删除e(只删除第一个)
func (a *Array) RemoveElement(e interface{}) bool {
	find := a.Find(e)
	if find != -1 {
		a.Remove(find)
		return true
	}
	return false
}

func (a *Array) resize(newCapacity int) {
	newArr := make([]interface{}, newCapacity)
	for i := 0; i < a.size; i++ {
		newArr[i] = a.data[i]
	}
	a.data = newArr
}
