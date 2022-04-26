package main

import (
	"fmt"
)

func a() {
	fmt.Println("Inside a()")
	defer func() {
		// 因此，当某个异常出现时，我们不会选择让解析器崩溃，
		// 而是会将panic异常当作普通的解析错误，并附加额外信息提醒用户报告此错误
		if c := recover(); c != nil {
			fmt.Println(c)
		}
	}()

	fmt.Println("About to call b()")
	b()
	fmt.Println("b() exited!")
	fmt.Println("Exiting a()")
}

func b() {
	fmt.Println("Inside b()")
	panic("Panic in b()")
	//fmt.Println("Exiting b()")
}

func main() {
	a()
	fmt.Println("main() ended!")
}
