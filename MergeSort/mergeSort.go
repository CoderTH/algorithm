package MergeSort

func Sort(arr []int) {
	var tamp = make([]int, len(arr))
	copy(tamp, arr)
	sort(arr, 0, len(arr)-1, tamp)
}
func sort(arr []int, l int, r int, tamp []int) {
	if l >= r {
		return
	}
	mid := l + (r-l)/2
	sort(arr, l, mid, tamp)
	sort(arr, mid+1, r, tamp)
	//加上如下判断
	//优化：当数组本身为有序时，这里的时间复杂度将会变成O(n)
	if arr[mid] > arr[mid+1] {
		merge(arr, l, mid, r, tamp)
	}
}

//合并两个有序的区间arr[l,mid],arr[mid+1,r]
func merge(arr []int, l int, mid int, r int, tamp []int) {
	copy(tamp[l:r+1], arr[l:r+1])
	i := l
	j := mid + 1
	//每轮循环为arr[k]赋值
	for k := l; k <= r; k++ {
		if i > mid {
			arr[k] = tamp[j]
			j++
		} else if j > r {
			arr[k] = tamp[i]
			i++
		} else if tamp[i] < tamp[j] {
			arr[k] = tamp[i]
			i++
		} else {
			arr[k] = tamp[j]
			j++
		}
		//arr[i]和arr[j]
	}
}
