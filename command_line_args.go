package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

/*
	Check for exactly two arguments
*/
func main() {
	if len(os.Args) == 1 {
		io.WriteString(os.Stderr, "Please give one or more arguments\n")
		os.Exit(1)
	} else if len(os.Args) != 3 {
		io.WriteString(os.Stderr, "Two arguments expected\n")
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

	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}
