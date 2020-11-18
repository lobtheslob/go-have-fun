package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func main() {

	stringChan := make(chan string)
	tower1Chan := make(chan string)
	tower2Chan := make(chan string)

	var offset int32
	offset = 3

	go tower1(stringChan, tower1Chan, offset)
	go tower2(stringChan, tower2Chan, offset)
	
	for i := 0, i < 2; i++; {
		select {
		case msg: <- tower1Chan:
			fmt.Printf("\n Control Tower: Message from Tower 1 - %v, msg")
		case msg: <- tower2Chan:
			fmt.Printf("\n Control Tower: Message from Tower 1 - %v, msg")
		}
	} 

}

func tower1(s chan string, t1 chan string, offset int32) {
	inputstream := bufio.NewReader(os.Stdin)
	fmt.Printf("Tower 1: Enter your message for Tower 2: ")
	userInput, _ := inputstream.ReadString('\n')
	userInput = strings.Replace(userInput,"\r\n", "", -1)
	
	var secretString string

	for _, c := range userInput {
		secretString += string(c+offset)
	}

	fmt.Printf("\n Tower 1: EncryptedString: %s", secretString)
	
	s <- secretString
	t1 <- "Msg sent to Tower 2"

}

func tower2(s chan string, t2 chan string, offset int32) {
	secretString := <- s 
	
	var orgString string

	for _, c := range secretString {
		orgString += string(c-offset)
	}	
	
	fmt.Printf("\nTower2: Decrpyted Message: %s: ", orgString)
	t2 <- "Msg received from Tower 1"
}