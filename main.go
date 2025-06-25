package main

import (
	// "bufio"								
	"fmt"
	// "os"
	"strings"
)

func wordcount(s string,ch chan int){

	words := strings.Fields(s)

	ch<- len(words)

}
func main(){


	sentences := []string{
   "Go was developed at Google to improve programming productivity with a focus on simplicity and fast compilation.",
		"Goroutines in Go provide an easy way to implement concurrency without the complexity of traditional threads.",
		"Go’s garbage collector and scheduler are optimized to deliver low-latency performance under heavy loads.",
		"Using channels in Go allows safe data sharing between goroutines without manual locking mechanisms.",
		"Key Go features like defer, structs, and interfaces help developers write clean and efficient code.",
}


	ch := make(chan int)

	for _ ,sentence := range sentences{
		go wordcount(sentence,ch)
	}

	sum:=0
	for i:=0 ;i<len(sentences);i++{
		words := <- ch
		sum += words
		fmt.Println("Words in respective sentence:", words)
	}

	fmt.Println("Total words:",sum)

}
