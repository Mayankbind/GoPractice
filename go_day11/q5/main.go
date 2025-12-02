package main

import (
	"fmt"
	"time"
)

func main(){
	channel1 := make(chan bool)
	fmt.Println("Task begin")
	go chanSend(channel1)
	go chanRecv(channel1)
	time.Sleep(time.Second*5)
	fmt.Println("Task completed")

}

func chanRecv(ch chan bool){
	fmt.Println(<- ch)
}

func chanSend(ch chan bool){
	ch <- true
}
