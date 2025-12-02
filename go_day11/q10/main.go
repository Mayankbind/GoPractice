package main

import (
	"errors"
	"fmt"
)


func main(){

	var err = function()

	if err != nil{
		fmt.Println("not finished")
		}else{
		fmt.Println("finished")
	}
}

func function() error{
	var err error
	if r := recover(); r != nil{
		err = errors.New("some error occured")
	}
	panic("panic")
	return err
}