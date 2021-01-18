package main

import "fmt"

func main() {
	var intSlice = []int{1, 2, 3, 4, 5}

	fmt.Print(intSlice)

	intSlice = append(intSlice, 6)

	partSlice := intSlice[1:3]

	fmt.Print(intSlice)

	fmt.Print(intSlice[1:3])

	fmt.Println(partSlice)

}
