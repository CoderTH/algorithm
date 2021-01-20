package test

import (
	"algorithm/SelectIonSort"
	"algorithm/SortingHelper"
	"testing"
)

func TestMergeSort(t *testing.T)  {
	n := 100000
	array := SelectIonSort.GenerateRandomArray(n, n)
	array2 := SelectIonSort.GenerateRandomArray(n, n)
	array3 := SelectIonSort.GenerateRandomArray(n, n)
	SortingHelper.SortTest("InsertionSort",array3,t)
	SortingHelper.SortTest("SelectionSort",array2,t)
	SortingHelper.SortTest("MergeSort",array,t)

}
