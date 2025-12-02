package main

import "fmt"

type Student struct{
	name string
}

func main(){
	defer func(){
		if r := recover(); r != nil{
			fmt.Println("Not finished")
		}
	}()
	mapper := make(map[int]struct{})

	mapper[0] = struct{}{}
	mapper[1] = struct{}{}
	mapper[2] = struct{}{}
	mapper[3] = struct{}{}
	mapper[4] = struct{}{}
	mapper[5] = struct{}{}


	for key := range mapper{
		mapper[key + 1] = struct{}{}
		delete(mapper, key)
	}

	// fmt.Println(mapper)

	fmt.Println("finished")
}
