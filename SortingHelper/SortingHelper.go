package SortingHelper

import (
	"algorithm/HeapSort"
	"algorithm/InsertionSort"
	"algorithm/MergeSort"
	"algorithm/QuickSort"
	"algorithm/SelectIonSort"
	"algorithm/ShellSort"
	"errors"
	"fmt"
	"testing"
	"time"
)

func IsSorted(arr []int) error {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			fmt.Println(arr[i],"----",arr[i+1])
			fmt.Println(i,"----",i+1)
			return errors.New("this sort is error ")
		}

	}
	return nil
}

func SortTest(sortName string, arr []int, t *testing.T) {
	start := time.Now()
	if sortName == "SelectionSort" {
		SelectIonSort.SelectionSort(arr)
	} else if sortName == "InsertionSort" {
		InsertionSort.InsertionSort(arr)
	} else if sortName == "MergeSort" {
		MergeSort.Sort(arr)
	} else if sortName == "QuickSort" {
		QuickSort.Sort(arr)
	} else if sortName == "QuickSort2Ways" {
		QuickSort.Sort2ways(arr)
	} else if sortName == "QuickSort3Ways" {
		QuickSort.Sort3ways(arr)
	}else if sortName == "HeapSort" {
		HeapSort.Sort(arr)
	}else if sortName == "ShellSort" {
		ShellSort.ShellSort(arr)
	}
	elapsed := time.Since(start)
	err := IsSorted(arr)
	if err != nil {
		t.Error(sortName, " ", err)
	}
	fmt.Println(sortName+" n=", len(arr), " time of use：", elapsed)
}
