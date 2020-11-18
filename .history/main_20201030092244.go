package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//uncomment for in order to add 2 threads
	//runtime.GOMAXPROCS(2)

	var waitGrp sync.WaitGroup
	waitGrp.Add(2)

	go func() {
		defer waitGrp.Done()


		time.Sleep(5 * time.Second)
		fmt.Println("Hello, ")
	}()

	go func() {
		defer waitGrp.Done()

		fmt.Println("Pluralsight")
	}()

	waitGrp.Wait()
}