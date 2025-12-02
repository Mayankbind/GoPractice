package main

import (
	"fmt"
	"sync"
)

func main(){
	mapper := make(map[string]struct{})

	mapper["value1"] = struct{}{}
	mapper["value2"] = struct{}{}

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(1)

	go func (mapper1 map[string]struct{}, group *sync.WaitGroup)  {
		mapper1["value3"] = struct{}{}
		mapper1["value4"] = struct{}{}

		fmt.Println(fmt.Sprintf("mapper1 length is %v", len(mapper1)))
		
		group.Done()
			
	}(mapper, &waitGroup)
		
		waitGroup.Wait()
		fmt.Println(fmt.Sprintf("mapper length is %v", len(mapper)))

}


