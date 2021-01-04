package main

import "fmt"

func main() {
	var arr = new([20]int)
	for i:=0;i<len(arr);i++ {
		arr[i] = i
		fmt.Print(arr[i]," ")
	}
		fmt.Println()
	scores := [...]int{100,99,66}
	for i:=0;i<len(scores);i++ {
		fmt.Print(scores[i]," ")
	}
	fmt.Println()
	for _,v:=range scores{
		fmt.Print(v," ")
	}
	scores[0] = 98
	fmt.Println()
	for _,v:=range scores{
		fmt.Print(v," ")
	}
}
