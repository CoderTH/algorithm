package test

import (
	"algorithm/Stack"
	"fmt"
	"testing"
)

func TestArrayStack(t *testing.T) {
	stack := Stack.NewArrayStack(10)
	for i := 0; i < 5; i++ {
		stack.Push(i)
		fmt.Println(stack)
	}
	stack.Pop()
	fmt.Println(stack)
}
