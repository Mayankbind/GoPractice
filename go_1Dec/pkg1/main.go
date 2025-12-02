package main

import (
	"fmt"
	"os"
	"io"
)

type Abser interface{
	Abs() float64
}

type Vertex struct{
	X, Y float64
}

func (v *Vertex) Abs() float64{
	return v.X + v.Y
}

type MyFloat float64

func (m MyFloat) Abs() float64{
	return float64(m)
}

func main(){
	file, err := os.Open("/home/mayank-bind/Documents/go_practice/go_1Dec/pkg1/temp.txt")

	fmt.Println("Error: ",err)
	defer file.Close()

	// var r io.Reader
	// buf  := make([]byte, 2048)

	// r = file
	// fmt.Println(r)
	// v, err := r.Read(buf)
	
	// fmt.Println("Bytes read: ",v)
	// fmt.Println("Data: ",string(buf[:v]))

	mp, err := countLetters(file)

	for key, value := range mp{
		fmt.Println(key," : ",value)
	}
	




}

func countLetters(r io.Reader) (map[string]int, error) {
	buf := make([]byte, 2048)
	byteCount, err := r.Read(buf)

	mp := make(map[string]int)

	for _, v := range buf{
		s := string(v)
		if(s >= "A" && s <= "Z" || s >= "a" && s <= "z") {
			mp[string(v)]++
		}
	}

	fmt.Println("Bytes read : ", byteCount)

	return mp, err
}
