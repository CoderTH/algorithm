package ShellSort

func ShellSort(arr []int) {
	for dis := len(arr)/2; dis >= 1; dis = dis/2 {
		for i:=dis;i<len(arr);i++ {
			for j :=i;j>=dis;j=j-dis{
				if arr[j]<arr[j-dis] {
					arr[j],arr[j-dis] =arr[j-dis],arr[j]
				}
			}
		}
	}
}


func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
}