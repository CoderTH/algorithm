package test

import (
	"algorithm/SelectIonSort"
	"algorithm/SortingHelper"
	"testing"
)

func TestHeapSort(t *testing.T) {
	n := 1000000
	arr := SelectIonSort.GenerateRandomArray(n, n)
	arr2 := SelectIonSort.GenerateRandomArray(n, n)
	arr3 := SelectIonSort.GenerateRandomArray(n, n)
	arr4 := SelectIonSort.GenerateRandomArray(n, n)

	SortingHelper.SortTest("MergeSort", arr, t)
	SortingHelper.SortTest("QuickSort2Ways", arr2, t)
	SortingHelper.SortTest("QuickSort3Ways", arr3, t)
	SortingHelper.SortTest("HeapSort", arr4, t)

}
