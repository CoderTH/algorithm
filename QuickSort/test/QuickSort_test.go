package test

import (
	"algorithm/SelectIonSort"
	"algorithm/SortingHelper"
	"testing"
)

func TestQuickSort(t *testing.T) {
	n := 1000000
	arr := SelectIonSort.GenerateRandomArray(n, n)
	arr2 := SelectIonSort.GenerateRandomArray(n, n)
	SortingHelper.SortTest("MergeSort",arr,t)
	SortingHelper.SortTest("QuickSort",arr2,t)
	SortingHelper.SortTest("MergeSort",arr,t)
	SortingHelper.SortTest("QuickSort",arr2,t)
}
