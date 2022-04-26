package main

import (
	"fmt"
	"math/rand"
	"time"
)

func add(c chan int) {
	sum := 0

	t := time.NewTimer(time.Second)

	for {
		select {
		case input := <-c:
			fmt.Println("input...")
			sum = sum + input
		case <-t.C:
			c = nil
			fmt.Println("nil...")
			fmt.Println(sum)
		}
	}
}

func send(c chan int) {
	for {
		fmt.Println("rand...")
		c <- rand.Intn(10)
	}
}

func main() {
	c := make(chan int)
	go add(c)
	go send(c)

	time.Sleep(10 * time.Second)
}
