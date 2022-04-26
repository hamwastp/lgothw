package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"
)

func f1(t int) {
	c1 := context.Background()
	c1, cancel := context.WithCancel(c1)
	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Need a delay!")
		return
	}

	delay, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Delay:", delay)

	f1(delay)
	// f2(delay)
	// f3(delay)
}
