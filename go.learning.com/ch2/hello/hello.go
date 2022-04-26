package main

/*
Relative paths don't work in the build context, because the build happens in a different directory from your source files.

You have a few choices to provide absolute paths:

You can use absolute paths in your source
You can use pkg-config to provide absolute paths
You can use the CGO_CFLAGS and CGO_LDFLAGS environment variables
You can use the ${SRCDIR} variable in the #cgo lines in your source.
See the cgo documentation for more details
*/

// #cgo CFLAGS: -I.
// #cgo LDFLAGS: ${SRCDIR}/hello.a -lstdc++
// #include <hello.h>
import "C"

//import "C" 需要空一行
func main() {
	C.SayHello(C.CString("Hello World\n"))
}
