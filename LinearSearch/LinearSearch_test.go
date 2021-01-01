package LinearSearch

import (
	"fmt"
	"testing"
)

func TestLinearSearch(t *testing.T) {

	//因为在golang中[]T无法直接转换成[]interface中，所以我们需要定义一个interface切片，将目标数组
	//中的数据写进去，然后再调用Search
	//说实话有点麻烦，但是我暂时没有找到好的解决方法
	tamp := []int{24,18,79,16,66,32,4,23}
	data := make([]interface{},len(tamp))
	for i,v:=range tamp{
		data[i] = v
	}
	//找到测试
	res := Search(data, 66)
	fmt.Println(res)
	//找不到测试
	res2 := Search(data, 99)
	fmt.Println(res2)
}
