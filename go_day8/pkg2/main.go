package main

import (
	"fmt"
	// "time"
	// "time"
)

func main() {
	var ch chan int
	m := make(map[string]int)
	fmt.Println(m)

	fmt.Printf("Channel Value: %v\n", ch)

	fmt.Println("Writing To the channel")
	// go TestNilCh(ch)
	// time.Sleep(1*time.Second)
	// fmt.Println("UNREACHABEL")

	
}

func TestNilCh(ch chan int){
	ch <-23
	fmt.Println("goroutine")
}