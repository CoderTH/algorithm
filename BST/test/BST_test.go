package test

import (
	"algorithm/Array/Array"
	"algorithm/BST"
	"fmt"
	"math/rand"
	"testing"
)

func TestBSt(t *testing.T) {
	bst := BST.NewBST()
	n := 1000
	for i := 0; i < n; i++ {
		bst.Add(rand.Intn(10000))
	}
	array := Array.NewArray(1)
	for !bst.IsEmpty(){
		array.AddLast(bst.RemoveMin())
	}
	fmt.Println(array)
	for i := 1; i < array.GetSize(); i++ {
		if array.Get(i-1).(int)>array.Get(i).(int) {
			panic("Error")
		}
	}
	fmt.Println("removeMin test completed")


	for i := 0; i < n; i++ {
		bst.Add(rand.Intn(10000))
	}
	array = Array.NewArray(1)
	for !bst.IsEmpty(){
		array.AddLast(bst.RemoveMax())
	}
	fmt.Println(array)
	for i := 1; i < array.GetSize(); i++ {
		if array.Get(i-1).(int)<array.Get(i).(int) {
			panic("Error")
		}
	}
	fmt.Println("removeMax test completed")

}
