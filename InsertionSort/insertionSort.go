package InsertionSort

//如果数组为有序的话，他的复杂度会变成O(n)
//插入排序优化
func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
}

func InsertionSort2(arr []int, l int, r int) {
	for i := l + 1; i <= r; i++ {
		for j := i; j > l && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
}
