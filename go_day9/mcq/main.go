package main

import (
	"fmt"
)

func main(){
	m := make(map[int]int)
	m[1] = 10
	m[2] = 20
	m[3] = 30
	m[4] = 40
	m[5] = 50
	// delete(m, 2)
	// fmt.Println(m)
	// fmt.Println(len(m))

	m3 := map[int]int{}
	m3[1] = 10
	m4 := m3
	m4[1] = 20

	func(x map[int]int){
		x[2] = 30
	}(m4)
	
	fmt.Println(m3)
	
	
	m2 := make(map[int][]int)
	m2[1] = append(m2[1], 99)

}
