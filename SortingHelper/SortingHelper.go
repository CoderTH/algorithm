package SortingHelper

import (
	"algorithm/SelectIonSort"
	"errors"
	"fmt"
	"testing"
	"time"
)

func IsSorted(arr []int) error {
	for i:=0;i<len(arr)-1;i++{
		if arr[i]>arr[i+1] {
			return errors.New("this sort is error ")
		}
	}
	return nil
}

func SortTest(sortName string,arr []int,t *testing.T)  {
	start:= time.Now()
	if sortName == "SelectionSort" {
		SelectIonSort.SelectionSort(arr)
	}
	elapsed := time.Since(start)
	err := IsSorted(arr)
	if err != nil {
		t.Error(sortName," ",err)
	}
	fmt.Println(sortName+" time of use：", elapsed)
}