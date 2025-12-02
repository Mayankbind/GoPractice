package main

import (
	"fmt"
)


type I interface{
	M()
}

type T struct{
	S string
}

type F float64

func (t *T) M(){
	if(t == nil){
		fmt.Println("nil")
		return
	}
	fmt.Println(t.S)
}

func (f F) M(){
	fmt.Println(f)
}

func do(i interface{}){
	switch v := i.(type){
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

type Person struct{
	Name string
	Age int
}

func (p Person) String() string{
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main(){
	p1 := Person{"Mayank", 12}
	p2 := Person{"Bind", 13}
	fmt.Println(p1, p2)
}

func describe(i interface{}){
	fmt.Printf("value : %v :: type : %T\n",i ,i)
}
