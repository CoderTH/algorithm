package MergeSort

func Sort(arr []int)  {
	sort(arr,0,len(arr)-1)
}
func sort(arr []int,l int,r int){
	if l >= r {
		return
	}
	mid := l+(r-l)/2
	sort(arr,l,mid)
	sort(arr,mid+1,r)
	merge(arr,l,mid,r)
}
//合并两个有序的区间arr[l,mid],arr[mid+1,r]
func merge(arr []int,l int,mid int,r int)  {
	var tamp  = make([]int,r-l+1)
	copy(tamp, arr[l:r+1])
	i:=l
	j:=mid+1
	//每轮循环为arr[k]赋值
	for k:=l;k<=r;k++{
		if i > mid {
			arr[k]=tamp[j-l]
			j++
		}else if j>r {
			arr[k] = tamp[i-l]
			i++
		}else if tamp[i-l]<tamp[j-l]{
			arr[k]=tamp[i-l]
			i++
		}else {
			arr[k]=tamp[j-l]
			j++
		}
		//arr[i]和arr[j]
	}
}
