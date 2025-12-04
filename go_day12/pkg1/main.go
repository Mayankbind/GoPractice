package main

import (
	"fmt"
	// "time"
)

func leakyGor(done <- chan bool)  {
		for{
			select {
			case <- done:
				return
			default :
			fmt.Println("Doing some work")
		}
	}
}


func sliceToChannel(nums []int) <- chan int{
	out := make(chan int)

	go func() {
		for _,val := range nums{
		out <- val
	}
	close(out)
	}()
	return out
}

func square(in <- chan int) <- chan int{
	out := make(chan int)

	go func() {
		for val := range in{
		out <- val * val
	}
	close(out)
	}()

	return out
}



func main(){
	nums := []int{1, 2, 3, 4}

	dataChannel1 := sliceToChannel(nums)

	dataChannel2 := square(dataChannel1)

	n := dataChannel2

	for val := range n {
		fmt.Println(val)
	}

}
