package main

import "fmt"


type Data struct{
	X, Y float64
}

func (d Data) Show() (float64, float64){
	return d.X, d.Y
}

func (d *Data) Scale() {
	d.X *= 2 
	d.Y *= 2 
}

func main(){
	// d := Data{1,2}
	// d.Scale()
	// fmt.Println(d.Show())

	// names := [4]string{
	// 	"mayank",
	// 	"bind",
	// 	"hello",
	// 	"world",
	// }

	// fmt.Println(names)
	
	// a := names[0:2]
	// b := names[1:3]
	
	// fmt.Println(a)
	// fmt.Println(b)
	
	// b[0] = "XXX"
	
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(names)

	array := [6]int{1,2,3}
	fmt.Println(array)

	b := make([]int, 5)
	fmt.Println(b[0])

}