package main

import "fmt"

func main() {
	i := -10
	j := 25

	pI := &i
	pJ := &j

	fmt.Println("pI memory:", pI)
	fmt.Println("pJ memory:", pJ)
	fmt.Println("pI value:", *pI)
	fmt.Println("pJ value:", *pJ)
}

func getPointer(n *int) {
	*n = *n * *n

}

func returnPointer(n int) *int {
	v := n * n
	return &v
}
