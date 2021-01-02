package LinearSearch
//数组生成器
func GenerateOrderedArray(n int)[]int  {
	arr := make([]int,n,2*n)
	for i:=0;i<len(arr);i++ {
		arr[i] = i
	}
	return arr
}