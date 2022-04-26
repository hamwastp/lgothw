package main

// #include <stdio.h>
import "C" // 启用CGO特性

func main() {
	C.puts(C.CString("Hello, world\n"))
}
