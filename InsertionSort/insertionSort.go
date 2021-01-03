package InsertionSort
//如果数组为有序的话，他的复杂度会变成O(n)
//插入排序优化
func InsertionSort(arr []int)  {
	for i:=0;i<len(arr);i++{
		tamp:=arr[i]
		var j int
		for j=i; j -1>=0&&arr[j]<arr[j-1];j--{
				arr[j-1] = arr[j]
		}
		arr[j] = tamp
	}
}

