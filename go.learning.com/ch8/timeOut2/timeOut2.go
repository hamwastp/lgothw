package main

import (
	"fmt"
	"sync"
	"time"
)

func timeout(w *sync.WaitGroup, t time.Duration) bool {
	temp := make(chan int)
	go func() {
		defer close(temp)
		time.Sleep(5 * time.Second)
		fmt.Println("=====0===")
		w.Done()
	}()

	select {
	case <-temp:
		fmt.Println("=====1===")
		return false
	case <-time.After(t):
		fmt.Println("=====2===")
		return true
	}

	// select 阻塞让出cpu， 以下代码不会执行到，
	//fmt.Println("=====3===")
	//return true
}

func main() {
	var w sync.WaitGroup
	w.Add(1)
	duration := time.Duration(3000) * time.Millisecond
	if timeout(&w, duration) {
		fmt.Println("Time out")
	} else {
		fmt.Println("ok")
	}
	w.Wait()
	fmt.Println("=====================================")
}
