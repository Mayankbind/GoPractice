package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)


func main(){
	var wg sync.WaitGroup

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	generator := func (dataItems string, stream chan interface{})  {
		for{
			select{
			case  <-ctx.Done():
				return

			case stream<-dataItems:
			}
		}
	}


	infiniteApple := make(chan interface{})
	go generator("apple", infiniteApple)

	infinitePeaches := make(chan interface{})
	go generator("peach", infinitePeaches)

	infiniteOranges := make(chan interface{})
	go generator("orange", infiniteOranges)


	wg.Add(1)
	go func1(&wg, ctx, infiniteApple)

	func2 := generic
	func3 := generic
	
	wg.Add(1)
	go func2(ctx, &wg, infiniteOranges)
	
	wg.Add(1)
	go func3(ctx, &wg, infinitePeaches)

	wg.Wait()
}

func func1(parentWaitGroup *sync.WaitGroup, ctx context.Context, infiniteApple chan interface{}){
	defer parentWaitGroup.Done()
	var wg sync.WaitGroup

	doWork := func(ctx context.Context){
		defer wg.Done()

		for{
			select{
			case <-ctx.Done():
				return

			case d, ok := <- infiniteApple :

				if !ok {
					fmt.Println("channel closed!!!")
				}
				fmt.Println(d)
				
			}
		}

	}

	newCtx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	for i := 0; i < 3; i++ {
		wg.Add(1)	
		go doWork(newCtx)
	}
	wg.Wait()
}

func generic(ctx context.Context, parentWaitGroup *sync.WaitGroup, stream <- chan interface{}){
	defer parentWaitGroup.Done()

	for{
		select{
		case <-ctx.Done():
			return
		case d, ok := <-stream:
			if !ok {
				fmt.Println("channel closed")
			}
			fmt.Println(d)

		}
	}

}
