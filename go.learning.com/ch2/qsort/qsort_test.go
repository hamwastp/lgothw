package qsort

// 单元测试不只是cgo

import (
	"testing"
	//extern int go_qsort_compare(void* a, void* b);
)

// #cgo CFLAGS: -I.
// import "C"

// //export go_qsort_compare
// func go_qsort_compare(a, b unsafe.Pointer) C.int {
// 	pa, pb := (*C.int)(a), (*C.int)(b)
// 	return C.int(*pa - *pb)
// }

func TestQsort(t *testing.T) {
	// values := []int32{42, 9, 101, 95, 27, 25}

	// qsort.Sort(unsafe.Pointer(&values[0]),
	// 	len(values), int(unsafe.Sizeof(values[0])),
	// 	qsort.CompareFunc(C.go_qsort_compare),
	// )
	//fmt.Println(values)
	//fmt.Println("h")
}
