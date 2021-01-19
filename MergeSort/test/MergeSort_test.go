package test

import (
	"algorithm/SelectIonSort"
	"algorithm/SortingHelper"
	"testing"
)

func TestMergeSort(t *testing.T)  {
	n := 100000
	array := SelectIonSort.GenerateRandomArray(n, n)
	SortingHelper.SortTest("MergeSort",array,t)
}
