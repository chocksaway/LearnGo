package main

import "fmt"

func main() {
	var counter int
	counter = 0

	for counter < 10 {
		counter++
		fmt.Println(counter)
	}

	fmt.Println("Exited loop counter = ", counter)
}
