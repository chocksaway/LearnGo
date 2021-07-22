package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	// Add a wait group
	var wg sync.WaitGroup
	for i := 1; i <= 1024; i++ {

		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				//fmt.Println(err) // this will print connection refused
				return
			}
			err = conn.Close()
			if err != nil {
				return
			}
			fmt.Printf("%d open\n", j)
		}(i)
	}
	wg.Wait() // block until wait group counter has returned to zero
}
