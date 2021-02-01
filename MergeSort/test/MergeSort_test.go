package test

import (
	"algorithm/SelectIonSort"
	"algorithm/SortingHelper"
	"testing"
)

func TestMergeSort(t *testing.T) {
	m := 100000
	array := SelectIonSort.GenerateRandomArray(m, m)
	array2 := SelectIonSort.GenerateRandomArray(m, m)
	array3 := SelectIonSort.GenerateRandomArray(m, m)
	SortingHelper.SortTest("InsertionSort", array3, t)
	SortingHelper.SortTest("SelectionSort", array2, t)
	SortingHelper.SortTest("MergeSort", array, t)

}
