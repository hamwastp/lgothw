package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func handleSingle(signal os.Signal) {
	fmt.Println("handleSingle() Caught:", signal)
}

func main() {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, os.Interrupt, syscall.SIGINT)

	go func() {
		for {
			sig := <-sigs
			switch sig {
			case os.Interrupt:
				fmt.Println("Caught:", sig)
			case syscall.SIGINT:
				handleSingle(sig)
				return
			}
		}
	}()

	for {
		fmt.Printf(".")
		time.Sleep(20 * time.Second)
	}
}
