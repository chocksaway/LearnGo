package main

import (
	"fmt"
	"sort"
)

type aStructure struct {
	person string
	height int
	weight int
}

func main() {
	var mySlice = make([]aStructure, 0)
	mySlice = append(mySlice, aStructure{"Mihalis", 180, 90})
	mySlice = append(mySlice, aStructure{"Bill", 134, 45})
	mySlice = append(mySlice, aStructure{"Marietta", 155, 45})
	mySlice = append(mySlice, aStructure{"Epifanios", 144, 50})
	mySlice = append(mySlice, aStructure{"Athina", 134, 40})

	fmt.Println("0:", mySlice)

	// height ascending order
	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].height < mySlice[j].height
	})

	fmt.Println("height ascending:", mySlice)

	// weight descending order
	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].weight > mySlice[j].weight
	})

	fmt.Println("weight descending:", mySlice)

}
