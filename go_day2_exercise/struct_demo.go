package main

import(
	"fmt"
)

//Struct in Go
	type Student struct {
		Name  string
		Age   int
		MathsScore int
		EnglishScore int
		ScienceScore int
	}


	func (s Student) averageScore() float64 {
		total := s.MathsScore + s.EnglishScore + s.ScienceScore
		average := float64(total) / 3.0
		return average
	}


	func main() {		
		s1 := Student{"Alice", 16, 85, 90, 78}
		avg := s1.averageScore()
		fmt.Printf("Average Score of %s: %.2f\n", s1.Name, avg)
	}