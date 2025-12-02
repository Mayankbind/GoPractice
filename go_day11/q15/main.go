package main

import "fmt"

func main(){
	mapper := make(map[int]struct{})
	prepare(mapper)
	fmt.Println(fmt.Sprintf("Values are %v", mapper))
}

func prepare(mapper map[int]struct{}){
	for index := 0; index <= 5; index++{
		defer func (){mapper[index] = struct{}{}}()
	}
}
