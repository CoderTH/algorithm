package test

import (
	"algorithm/SelectIonSort"
	"algorithm/SortingHelper"
	"testing"
)

func TestInsertSort(t *testing.T)  {
	dataSize := []int{10000,100000}
	for _,v:=range dataSize{
		arr := SelectIonSort.GenerateRandomArray(v,v)
		arr2 := make([]int,v,v)
		copy(arr2,arr)
		SortingHelper.SortTest("InsertionSort",arr,t)
		SortingHelper.SortTest("InsertionSort2",arr2,t)
	}
}
