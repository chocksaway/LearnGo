package main

import "fmt"

func main() {
	var aSliceLiteral = []int{1, 2, 3, 4, 5}

	fmt.Print(aSliceLiteral)

	aSliceLiteral = append(aSliceLiteral, 6)

	fmt.Print(aSliceLiteral)

}
