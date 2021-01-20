package main

import "fmt"

func main() {
	var intSlice = []int{1, 2, 3, 4, 5}

	fmt.Println(intSlice)

	fmt.Printf("Cap: %d, Length: %d\n", cap(intSlice), len(intSlice))

	intSlice = append(intSlice, 6)

	partSlice := intSlice[1:3]

	fmt.Println(intSlice)

	fmt.Printf("Cap: %d, Length: %d\n", cap(intSlice), len(intSlice))

	fmt.Print(intSlice[1:3])

	fmt.Println(partSlice)

}
