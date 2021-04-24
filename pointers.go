package main

import "fmt"

func main() {
	k := 100
	pK := &k

	fmt.Println("p:", k)
	fmt.Println("pK", *pK)

	*pK = 200

	fmt.Println("p:", k)
	fmt.Println("pK", *pK)
}
