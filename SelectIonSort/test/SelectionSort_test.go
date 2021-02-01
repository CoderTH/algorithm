package test

import (
	"algorithm/SelectIonSort"
	"algorithm/SortingHelper"
	"fmt"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	arr := []int{1, 4, 2, 3, 6, 5}
	SelectIonSort.SelectionSort(arr)
	for _, v := range arr {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func TestGenerateSort(t *testing.T) {
	dataSize := []int{10000, 100000}
	for _, v := range dataSize {
		arr := SelectIonSort.GenerateRandomArray(v, v)
		SortingHelper.SortTest("SelectionSort", arr, t)
	}
}
