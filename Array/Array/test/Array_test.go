package test

import (
	"algorithm/Array/Array"
	"fmt"
	"testing"
)

func TestArray(t *testing.T)  {
	arr := Array.NewArray(10)
	for i:=0;i<10;i++ {
		arr.AddLast(i)
	}
	fmt.Print(arr)
	arr.Add(1,100)
	arr.Add(1,"nihao")
	fmt.Print(arr)
	arr.AddFirst(-1)
	fmt.Print(arr)
	arr.Remove(1)
	fmt.Print(arr)
	arr.RemoveElement("nihao")
	fmt.Println(arr)
	arr.RemoveFirst()
	fmt.Println(arr)
}
