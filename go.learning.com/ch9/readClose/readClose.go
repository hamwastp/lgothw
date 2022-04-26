package main

import "fmt"

func main() {
	willClose := make(chan int, 10)

	willClose <- 1
	willClose <- 2
	willClose <- 3
	<-willClose
	<-willClose
	//<-willClose

	close(willClose)
	fmt.Println(<-willClose)
}
