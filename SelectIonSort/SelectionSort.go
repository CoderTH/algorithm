package SelectIonSort

//这里没办法做到像视频中使用范型得到的效果，golang中拥有一个sort.interface接口，如果是自定义类型，可以通过
//实现这个接口，然后通过效用less来判断大小，但是貌似golang中基本的数据类型并没有实现这个接口，所以没有办法实现适配所有类型的排序
//这里就只实现了int切片的排序，如果大家有解决方案可以联系我CoderTh@outlook.com
func SelectionSort(arr []int) {
	//维护一个循环不变量
	//arr[0...i)是有序的；arr[i...n]是无序的
	for i := 0; i < len(arr); i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			//if !arr[i].Less(i,minIndex){
			//	minIndex = j
			//}
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}
		swap(arr, i, minIndex)
	}
}

//这里简单介绍下golang中参数传值的机制
//在golang中其实函数的参数传值默认是值传递，但是类似于struct，map，切片这中类型，他们的传值
//其实默认是引用传递（有一些歧义）
//为什么说他拥有一些歧义呢，总结了以下几点
/*
对于普通类型(int, string 等等), 就是 传值 调用, 函数内对参数的修改, 不影响外面的变量
对于 struct 指针, slice 和 map 类型, 函数内对参数的修改之所以能影响外面, 是因为参数和外面的变量指向了同一块数据的地址
对于 struct 指针, slice 和 map 类型, 函数的参数和外面的变量的地址是不一样的, 所以本质上还是 传值 调用
slice 的 append 操作会改变 slice 指针的地址,
*/
//有兴趣的可以去查一下资料或者自己实践一下，我们这里着重算法，我们就简单的将其理解成引用传递
func swap(arr []int, i int, minIndex int) {
	arr[i], arr[minIndex] = arr[minIndex], arr[i]
}

//自定义类，可通过次方法排序
//func SelectionCustomSort(arr []sort.Interface)  {
//	//维护一个循环不变量
//	//arr[0...i)是有序的；arr[i...n]是无序的
//	for i:=0;i<len(arr);i++{
//
//		for j := i+1; j < len(arr); j++ {
//			if !arr[i].Less(i,j){
//				arr[i],arr[j] = arr[j],arr[i]
//			}
//		}
//	}
//}
