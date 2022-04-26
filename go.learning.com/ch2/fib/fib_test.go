package fib

import (
	"fmt"
	"testing"
)

func TestFibList(t *testing.T) {
	// var a int = 1
	// var b int = 1

	// var (
	// 	a int = 1
	// 	b int = 1
	// )

	a := 1 //类型推断
	b := 1 //类型推断

	for i := 1; i < 5; i++ {
		fmt.Print(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}

	fmt.Println()
}
