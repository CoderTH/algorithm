package QuickSort

import "math/rand"

func Sort(arr []int) {
	sort(arr,0,len(arr)-1)
}

func sort(arr []int,l int,r int)  {
	if l >= r {
		return
	}
	p := partition(arr, l, r)
	sort(arr,l,p-1)
	sort(arr,p+1,r)
}

func partition(arr []int,l int,r int) int {
	var p int = rand.Intn(r-l+1)+l
	arr[l],arr[p]=arr[p],arr[l]
	j := l
	for i:= l+1;i<=r;i++{
		if arr[i]<arr[l]{
			j++
			arr[i],arr[j] = arr[j],arr[i]
		}
	}
	arr[l],arr[j] = arr[j],arr[l]
	return j
}
