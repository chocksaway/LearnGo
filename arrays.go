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

	var twoD = [4][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12},
		{13, 14, 15, 16}}

	fmt.Println(len(twoD[0]))

	var counterInner int
	for counter = 0; counter < len(twoD); counter++ {
		for counterInner = 0; counterInner < len(twoD[counter]); counterInner++ {
			fmt.Print(twoD[counter][counterInner])
		}
	}

	fmt.Println("\nusing _")

	for _, outer := range twoD {
		for _, inner := range outer {
			fmt.Print(inner)
		}
	}

}
