package main

import (
	"fmt"
	// "runtime/debug"
	// "strings"
	// "time"
)

func doPanic(msg string){
	panic(msg)
}

func Temp(msg string){
	fmt.Println("Temp is called...",msg)
}

func div60(i int){
	defer func ()  {
		if v := recover(); v != nil {
			fmt.Println(v)
		}
	}()

	fmt.Println(60/i)
}

func Test(msg int){
	defer DeferTest(msg)
	DeferTest(msg)
	DeferTest(msg)
	defer DeferTest(msg)
	defer Temp("Main defer")
}

func DeferTest(msg int){
	fmt.Println(msg)
}

type Student struct{
	name string
	id int 
}

func SliceVal(s *[]int){
	(*s)[0] = 100
	*s = append(*s, 25)
	fmt.Println(len(*s))
	fmt.Println(cap(*s))
}

func chanTest(ch chan int){
	ch <- 10
	ch <- 20
}

func main(){

	ch := make(chan int)
	go chanTest(ch)

	for v := range ch {
		fmt.Println(v)
	}


	// s := make([]int, 3, 5)
	// SliceVal(&s)
	// fmt.Println(s)
	// fmt.Println(len(s))

	// for _, v := range []int{1, 2, 3, 0}{
	// 	div60(v)
	// }
	
	// defer func ()  {
	// 	fmt.Println(recover())
	// }()

	// defer func ()  {
	// 	panic("Inner Panic")
	// }()

	// panic("Outer Panic")

	
	
	defer func ()  {
		// if r := recover(); r != nil{
		// 	fmt.Println(r)
		// 	// fmt.Println(string(debug.Stack()))
		// }
		
		fmt.Println(recover())
		fmt.Println(recover())
		fmt.Println(recover())
		}()

		panic("Hello")

		// go DemoGor1()
		// go DemoGor2()
		// go DemoGor3()
		// go DemoGor4()
		// time.Sleep(1000*time.Millisecond)
		
}

func DemoGor1(){
	panic("Hello1")
}

func DemoGor2(){
	panic("Hello2")
}

func DemoGor3(){
	panic("Hello3")
}

func DemoGor4(){
	panic("Hello4")
}