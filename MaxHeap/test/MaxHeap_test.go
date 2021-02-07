package test

import (
	"algorithm/MaxHeap"
	"math/rand"
	"testing"
)

func TestMaxHeap(t *testing.T) {
	n := 1000000
	maxHeap := MaxHeap.NewMaxHeap(1)
	for i := 0; i < n; i++ {
		maxHeap.Add(rand.Intn(n))
	}
	var arr []int
	for i := 0; i < n; i++ {
		arr = append(arr, maxHeap.ExtractMax())
	}
	for i := 1; i < n; i++ {
		if arr[i-1] < arr[i] {
			panic("Error")
		}
	}
}
