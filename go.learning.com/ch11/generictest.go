package main

import (
	"fmt"
	"strconv"
)

func printSlice[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
	fmt.Print("\n")
}

func Must2[T1, T2 any](t1 T1, t2 T2, err error) (T1, T2) {
	if err != nil {
		panic(err)
	}
	return t1, t2
}

func foo() (int, float64, error) {
	return 1, 2, nil
}

// Vector 是一个切片，其元素的类型由类型参数 T 确定。
// T 的实际类型需要在声明 Vector 对象的时候指定。
type Vector[T any] []T

func (v *Vector[T]) Push(x T) {
	*v = append(*v, x)
}

type Stringer interface {
	String() string
}

func Stringify[T Stringer](s []T) (ret []string) {
	for _, v := range s {
		ret = append(ret, v.String()) //错误
	}

	return ret
}

type MyInt int

func (i MyInt) String() string { return strconv.Itoa(int(i)) }

func main() {
	printSlice[int]([]int{1, 2, 3, 4, 5})
	printSlice([]int{1, 2, 3, 4, 5})

	i, f := Must2[int, float64](foo())
	println(i, f)

	var v Vector[int] // v的类型为[]int
	v.Push(1)
	println(v[0])

	Stringify([]MyInt{1, 2, 3}) // 返回 []string{"1","2","3"}
}
