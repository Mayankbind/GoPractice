package main

import "fmt"


func sum(s []int, c chan int){
	sum := 0
	for _,v := range s {
		sum += v
	}
	c<-sum
}

func main(){
	c := make(chan int, 2)
	s := []int{1, 2, 3, 4, 5, 6}
	go sum(s[len(s)/2:], c)
	go sum(s[:len(s)/2], c)

	x, y := <-c, <-c
	fmt.Println("x: ",x)
	fmt.Println("y: ",y)

	bufferedChannel := make(chan int, 2)
	bufferedChannel<-1
	bufferedChannel<-2
	close(bufferedChannel)
	for i := range bufferedChannel{
		fmt.Println(i)
	}
}