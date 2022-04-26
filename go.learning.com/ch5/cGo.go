package main

//#include <stdio.h>
//void callC() {
//	  printf("calling c code!\n");
//}
import "C"
import (
	"fmt"
	"sort"
)

func main() {
	C.callC()

	aMap := map[string]int{}
	aMap["test"] = 1
	aMap = nil
	fmt.Println(aMap)

	mySlice := make([]aStructure, 0)
	mySlice = append(mySlice, aStructure{"Mihailis", 180, 90})
	mySlice = append(mySlice, aStructure{"Bill", 134, 45})
	mySlice = append(mySlice, aStructure{"Bill1", 135, 45})

	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].height < mySlice[j].height
	})

	fmt.Println(">:", mySlice)
}

type aStructure struct {
	person string
	height int
	weight int
}
