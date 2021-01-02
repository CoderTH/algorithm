package SelectIonSort

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestSelectionSort(t *testing.T) {
	arr := []int{1,4,2,3,6,5}
	SelectionSort(arr)
	for _,v:=range arr{
		fmt.Print(v," ")
	}
	fmt.Println()
}

func TestGenerateSort(t *testing.T) {
	arr := GenerateRandomArray(10000,10000)
	start:= time.Now()
	SelectionSort(arr)
	elapsed := time.Since(start)
	fmt.Println(" SelectionSort time of use：", elapsed)
	err := CheckSort(arr)
	if err != nil {
		t.Error(err)
	}else {
		fmt.Println("Sort check pass")
	}
}

func CheckSort(arr []int) error {
	for i:=0;i<len(arr)-1;i++{
		if arr[i]>arr[i+1] {
			return errors.New("this sort is error ")
		}
	}
	return nil
}
