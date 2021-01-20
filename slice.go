package main

import "fmt"

func main() {
	var intSlice = []int{1, 2, 3, 4, 5}
	var intSlice2 = []int{-1, 6}

	// destination source
	copy(intSlice, intSlice2)

	fmt.Println(intSlice)

	fmt.Println(intSlice2)

	fmt.Printf("Cap: %d, Length: %d\n", cap(intSlice), len(intSlice))

	intSlice = append(intSlice, 6)

	partSlice := intSlice[1:3]

	fmt.Println(intSlice)

	fmt.Printf("Cap: %d, Length: %d\n", cap(intSlice), len(intSlice))

	fmt.Print(intSlice[1:3])

	fmt.Println(partSlice)

}
