package main

import (
	"fmt"
	// "time"
)


func mapAdd(mp map[int]string){
	for i := 0; i < 10; i++{
		mp[i+1] = "G1"
	} 
}

func mapAdd2(mp map[int]string){
	for i := 0; i < 10; i++{
		mp[i+1] = "G2"
	} 
}

func main(){
	// Zero maps 
	var mp map[string]int

	fmt.Println(mp == nil) //true
	
	
	// mapAdd(mp) // throws error
	fmt.Println(mp["temp"]) //true


	mp2 := make(map[int]string)

	// go mapAdd(mp2)
	// go mapAdd2(mp2)
	// time.Sleep(1*time.Second)

	showMap(mp2)

	//Using slice as map key by serialize struct to string 
	mp3 := make(map[string]int)

	k := Key{
		ID : 1,
		Tags: []string{"hello", "world"},
	}

	mp3[k.String()] = 2
	fmt.Println(mp3)

	
}

type Key struct{
	ID int
	Tags []string
}

func (k Key) String() string{
	return fmt.Sprintf("%v : %v", k.ID, k.Tags)
}

func showMap(mp map[int]string){
	for i := 0; i < len(mp); i++{
		fmt.Println(i+1," : ",mp[i+1])
	} 

}
	

