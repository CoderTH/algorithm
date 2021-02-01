package test

import (
	"algorithm/SelectIonSort"
	"algorithm/SortingHelper"
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	n := 100000
	arr := SelectIonSort.GenerateRandomArray(n, n)
	arr2 := SelectIonSort.GenerateRandomArray(n, n)
	arr3 := SelectIonSort.GenerateRandomArray(n, n)
	SortingHelper.SortTest("QuickSort", arr, t)
	SortingHelper.SortTest("QuickSort2Ways", arr2, t)
	SortingHelper.SortTest("QuickSort3Ways", arr3, t)
	fmt.Println()
	arr = SelectIonSort.GenerateOrderedArray(n)
	arr2 = SelectIonSort.GenerateOrderedArray(n)
	arr3 = SelectIonSort.GenerateOrderedArray(n)
	SortingHelper.SortTest("QuickSort", arr, t)
	SortingHelper.SortTest("QuickSort2Ways", arr2, t)
	SortingHelper.SortTest("QuickSort3Ways", arr3, t)
	fmt.Println()
	n = 50000
	arr = SelectIonSort.GenerateRandomArray(n, 1)
	arr2 = SelectIonSort.GenerateRandomArray(n, 1)
	arr3 = SelectIonSort.GenerateRandomArray(n, 1)
	SortingHelper.SortTest("QuickSort", arr, t)
	SortingHelper.SortTest("QuickSort2Ways", arr2, t)
	SortingHelper.SortTest("QuickSort3Ways", arr3, t)
}
