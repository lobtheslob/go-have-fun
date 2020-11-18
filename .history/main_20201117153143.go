package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//uncomment for in order to add 2 threads
	//runtime.GOMAXPROCS(2)

//	var waitGrp sync.WaitGroup
//	waitGrp.Add(2)
//
//	go func() {
//		defer waitGrp.Done()
//
//
//		time.Sleep(5 * time.Second)
//		fmt.Println("Hello, ")
//	}()
//
//	go func() {
//		defer waitGrp.Done()
//
//		fmt.Println("Pluralsight")
//	}()
//
//	waitGrp.Wait()

//rand.Seed(time.Now().UnixNano())

infiniteloop(rand.Intn(2))

}

// find out if a randnumber that someone passes is 
//the same as the method's randnumber are the same
func infiniteloop(randNum int) {

	count := 1;
	done := false;

	for {
	ourNum := rand.Intn(1000)
		if (randNum == ourNum) {
			done = true;
			break;
		}
		count++
	}

   	if (done) {
   	    fmt.Printf("Number #%v was found after %v attempts", randNum, count)
   	}	
}
