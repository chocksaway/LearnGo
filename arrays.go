package main

import "fmt"

func main() {
	var counter int
	counter = 0

	var anArray = [4]int{1, 2, 4, -4}

	for counter < len(anArray) {
		fmt.Println(anArray[counter])
		counter++
	}

	fmt.Println("Exited loop")
}
