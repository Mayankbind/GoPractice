package main

import "fmt"

func main(){
	function()
	fmt.Println("finished")
}

func function(){
	if r := recover(); r != nil{
		fmt.Println("panic recovered")
	}
	
	fmt.Println("some work-1")
	panic("panic")
	fmt.Println("some work-2")
}
