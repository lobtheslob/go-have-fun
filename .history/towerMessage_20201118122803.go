package main

func main() {

	stringChan := make(chan string)
	tower1Chan := make(chan string)
	tower2Chan := make(chan string)

	go tower1(stringChan, tower1Chan, offset)
	go tower2(stringChan, tower2Chan, offset)

}

tower1(s chan string, t1 chan string, offset int32) {
	inputstream := bufio.NewReader(os.Stdin)
	fmt.Printf()
}