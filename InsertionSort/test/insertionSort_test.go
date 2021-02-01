package test

import (
	"algorithm/SelectIonSort"
	"algorithm/SortingHelper"
	"fmt"
	"testing"
)

func TestInsertSort(t *testing.T) {
	dataSize := []int{10000, 100000}
	for _, v := range dataSize {
		fmt.Println("Random Array : ")
		arr := SelectIonSort.GenerateRandomArray(v, v)
		arr2 := make([]int, v, v)
		copy(arr2, arr)
		SortingHelper.SortTest("InsertionSort", arr, t)
		SortingHelper.SortTest("SelectionSort", arr2, t)

		fmt.Println("Ordered Array : ")
		arr = SelectIonSort.GenerateOrderedArray(v)
		copy(arr2, arr)
		SortingHelper.SortTest("InsertionSort", arr, t)
		SortingHelper.SortTest("SelectionSort", arr2, t)
	}
}
