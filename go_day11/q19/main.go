package main

import (
	"errors"
	"fmt"
)

func main(){
	var err = function()
	fmt.Println(fmt.Sprintf("error %v", err))

	var temp = TempFun()
	fmt.Println(temp)

}
 func function() (err1 error){
	defer func ()  {
		if r := recover(); r != nil{
			err1 = errors.New("some error occured")
		} 
	}()
	panic("panic")
	return err1
 }


 func TempFun() (x int){
	return
}
