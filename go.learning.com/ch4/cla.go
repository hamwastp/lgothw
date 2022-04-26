package main

import (
	"fmt"
	"os"
	"strconv"
)

func main1() {
	if len(os.Args) == 1 {
		fmt.Println("Please give one or more floats")
		os.Exit(1)
	}

	arguments := os.Args
	min, _ := strconv.ParseFloat(arguments[1], 64)
	max, _ := strconv.ParseFloat(arguments[1], 64)

	for i := 2; i < len(arguments); i++ {
		n, _ := strconv.ParseFloat(arguments[i], 64)

		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	fmt.Println("min: ", min)
	fmt.Println("max: ", max)
}
