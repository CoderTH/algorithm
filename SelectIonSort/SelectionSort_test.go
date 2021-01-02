package SelectIonSort

import (
	"fmt"
	"testing"
)


func TestSelectionSort(t *testing.T) {
	arr := []int{1,4,2,3,6,5}
	SelectionSort(arr)
	for _,v:=range arr{
		fmt.Print(v," ")
	}
	fmt.Println()
}
