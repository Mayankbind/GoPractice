package main

import (
	"context"
	"fmt"
)

func main(){
	ctx := context.Background()
	ctx1, cancel1 := context.WithCancel(ctx)
	ctx2, cancel2 := context.WithCancel(ctx1)
	ctx3, cancel3 := context.WithCancel(ctx2)

	cancel2()

	fmt.Println(cancel1)
	fmt.Println(cancel3)
	fmt.Println(cancel2)
	
	fmt.Println(ctx3)
	fmt.Println(ctx2)
	fmt.Println(ctx1)
	fmt.Println(ctx)

}