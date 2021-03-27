package QuickSort

import "math/rand"

func Sort(arr []int) {
	sort(arr, 0, len(arr)-1)
}

func sort(arr []int, l int, r int) {
	if l >= r {
		return
	}
	p := partition(arr, l, r)
	sort(arr, l, p-1)
	sort(arr, p+1, r)
}

func partition(arr []int, l int, r int) int {
	var p int = rand.Intn(r-l+1) + l
	arr[l], arr[p] = arr[p], arr[l]//将标定点元素换到第一个
	j := l
	for i := l + 1; i <= r; i++ {
		if arr[i] < arr[l] {
			j++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[l], arr[j] = arr[j], arr[l]
	return j
}

func Sort2ways(arr []int) {
	sort2ways(arr, 0, len(arr)-1)
}

func sort2ways(arr []int, l int, r int) {
	if l >= r {
		return
	}
	p := partition2(arr, l, r)
	sort2ways(arr, l, p-1)
	sort2ways(arr, p+1, r)
}
func partition2(arr []int, l int, r int) int {
	var p int = rand.Intn(r-l+1) + l
	arr[l], arr[p] = arr[p], arr[l]
	//arr[l+1...i-1]<=v;arr[j+1...r]>=v
	i := l + 1
	j := r
	for {
		for i <= j && arr[i] < arr[l] {
			i++
		}
		for j >= i && arr[j] > arr[l] {
			j--
		}
		if i >= j {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	arr[l], arr[j] = arr[j], arr[l]
	return j
}

func Sort3ways(arr []int) {
	sort3ways(arr, 0, len(arr)-1)
}

func sort3ways(arr []int, l int, r int) {
	if l >= r {
		return
	}
	var p int = rand.Intn(r-l+1) + l
	arr[l], arr[p] = arr[p], arr[l]
	lt := l
	i := l + 1
	gt := r + 1
	for i < gt {
		if arr[i] < arr[l] {
			lt++
			arr[i], arr[lt] = arr[lt], arr[i]
			i++
		} else if arr[i] > arr[l] {
			gt--
			arr[i], arr[gt] = arr[gt], arr[i]
		} else {
			i++
		}
	}
	arr[l], arr[lt] = arr[lt], arr[l]
	sort3ways(arr, l, lt-1)
	sort3ways(arr, gt, r)
}
