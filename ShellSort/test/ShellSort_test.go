package test

import (
	"algorithm/SelectIonSort"
	"algorithm/SortingHelper"
	"fmt"
	"testing"
)

func TestShellSort(t *testing.T) {
	n := 10000
	arr := SelectIonSort.GenerateRandomArray(n, n)
	arr3 := SelectIonSort.GenerateRandomArray(n, n)
	SortingHelper.SortTest("ShellSort", arr, t)
	SortingHelper.SortTest("MergeSort", arr3, t)
	fmt.Println()
	arr = SelectIonSort.GenerateOrderedArray(n)
	arr3 = SelectIonSort.GenerateOrderedArray(n)
	SortingHelper.SortTest("ShellSort", arr, t)
	SortingHelper.SortTest("MergeSort", arr3, t)
	fmt.Println()
	n = 10000
	arr = SelectIonSort.GenerateRandomArray(n, 1)
	arr3 = SelectIonSort.GenerateRandomArray(n, 1)
	SortingHelper.SortTest("ShellSort", arr, t)
	SortingHelper.SortTest("MergeSort", arr3, t)

}
