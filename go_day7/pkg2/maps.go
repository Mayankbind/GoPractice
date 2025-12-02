package main

import (
	"fmt"
	"strings"
	// "strconv"
)

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Google": Vertex{1.23, 4.21},
	"Yahoo": Vertex{2.32, 4.53},
}

func add (fn func(x, y int)){
	fn(1, 2)
}

func main(){

	m = make(map[string]Vertex)
	m["google"] = Vertex{1.01, 2.02}
	_, ok := m["google"]
	fmt.Println(ok)

	s := "hello world temp"
	v := strings.Split(s, " ")
	fmt.Println(v[0] == "hello")

	hpot := func (x, y int)  {
		fmt.Println("Anonymous function",x ,y)
	}

	hpot(1, 2)
	add(hpot)

}
