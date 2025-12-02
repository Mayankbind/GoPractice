package main
import "fmt"

type Student struct{
	name string
	id int
}

func main() {
	s := Student{
		name: "Mayank",
		id: 24,
	}

	m := map[int]Student{
		1: s,
		2: Student{
			name: "Ankit",
			id: 25,
		},
	}

	temp := m[1]
	temp.id = 34
	m[1] = temp
	fmt.Println(m[1])
	fmt.Println(temp)
}