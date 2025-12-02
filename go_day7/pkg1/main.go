package main

import "fmt"

var pow = []int {1, 2, 4, 8, 16, 32, 64, 128}

func Show(s []int){
	for _, v := range s{
		fmt.Print(v)
	} 
	fmt.Println()
}

func main(){

	s := []int{1, 2, 3, 4, 5}
	
	Show(s)

	// t := s
	// s = append(s, t...)

	t := make([]int, 5)
	copy(t,s)
	t[0] = 100
	fmt.Println(t)
	fmt.Println(s)



}
