package SelectIonSort

import (
	"math/rand"
	"time"
)

//数组生成器
func GenerateOrderedArray(n int) []int {
	arr := make([]int, n, 2*n)
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}
	return arr
}

func GenerateRandomArray(n int, bound int) []int {
	arr := make([]int, 0, 0)
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		arr = append(arr, rand.Intn(bound))
	}
	return arr
}
