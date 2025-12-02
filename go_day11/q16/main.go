package main

import "fmt"

func main(){
	// var x int 
	// arr := [3]int{3, 5, 2}
	// // x, arr = arr[0], arr[1:]
	// arr = [3]int{1, 2}

	// fmt.Println(x, arr)
	// fmt.Printf("\n%T\n",arr)

	i := 0
	defer func(i int){
		fmt.Println(i)
	}(i)
	i++

}
