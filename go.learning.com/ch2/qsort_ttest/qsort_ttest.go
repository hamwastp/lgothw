package qsort_ttest

/*
extern int  _cgo_qsort_compare(void* a, void* b);
*/

// #cgo CFLAGS: -I.

import "C"

import (
	"fmt"
	"unsafe"

	"go.learning.com/ch2/qsort"
)

//export _cgo_qsort_compare
func _cgo_qsort_compare(a, b unsafe.Pointer) C.int {
	pa, pb := (*C.int)(a), (*C.int)(b)
	return C.int(*pa - *pb)
}

func main() {
	values := []int32{42, 9, 101, 95, 27, 25}

	qsort.Sort(unsafe.Pointer(&values[0]),
		len(values), int(unsafe.Sizeof(values[0])),
		qsort.CompareFunc(C._cgo_qsort_compare),
	)
	fmt.Println(values)
	//fmt.Println("h")
}
