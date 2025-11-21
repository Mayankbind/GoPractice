package main

import "fmt"

func show(x, y string) (a, b string) {
	a, b = x, y
	return y, x
}

func main() {
	temp := 10
	fmt.Println(show("Hello", "World"))
	fmt.Println("Temp value is:", temp)

}
