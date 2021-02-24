package main

import "fmt"

const const_str string = "milesd"

func main() {
	var anotherMap = map[string]int{
		"k1": 12,
		"k2": 13,
	}

	delete(anotherMap, "k2")

	fmt.Println(anotherMap)

	_, ok := anotherMap["Tom"]

	if ok {
		fmt.Println("Exists!")
	} else {
		fmt.Println("Does NOT exist")
	}

	anotherMap["k3"] = 22

	for key, value := range anotherMap {
		fmt.Println(key, value)
	}

	fmt.Print(const_str)
}
