package main

import(
	"fmt"
)

//Struct in Go
	type Student struct {
		Name  string
		Age   int
		Grade string
	}

	func (s *Student) displayInfo() {
		s.Name = "Changed Name"
		fmt.Println("Name:", s.Name)
	}

func main() {

	s1 := Student{"Mayank", 20, "A"}
	s1.displayInfo()
	fmt.Println("Name after method call:", s1.Name)

	// var a [3] int = [3]int{10,20,30}
	// var c = []int{100,200,300}
	// c = append(c,400)
	// a[1]=50

	// for i:=0; i< len(a); i++ {
		
	// 	fmt.Println(a[i])
	// }

	// fmt.Println("Array c:")
	// for i:=0; i< len(c); i++ {
	// 	fmt.Println(c[i])
	// }


	//Slice Capacity and Length
	// var s []int
	// fmt.Println("Length: %d, Capacity: %d\n", len(s), cap(s))
	
	// s = append(s, 10)
	// fmt.Println("Length: %d, Capacity: %d\n", len(s), cap(s))
	
	// s = append(s, 20, 30, 40)
	// fmt.Println("Length: %d, Capacity: %d\n", len(s), cap(s))

	// s = append(s, 50)
	// fmt.Println("Length: %d, Capacity: %d\n", len(s), cap(s))

	// s = append(s, 60, 70, 80, 90)
	// fmt.Println("Length: %d, Capacity: %d\n", len(s), cap(s))		



	// message := fmt.Sprintf("Hello %v your score is %v", "Mayank", 100)

	// fmt.Println(message)

	// m, n := add(5, 10)

	// fmt.Println("Values of m and n are:", m, n)


	// Function exercise
	// var revenue float64
	// var expenses float64
	// var taxRate float64

	// revenue = getUserInput("revenue")
	// expenses = getUserInput("expenses")
	// taxRate = getUserInput("tax rate")

	// ebt := revenue - expenses
	// profit := ebt * (1 - taxRate/100)
	// ratio := ebt / profit

	// fmt.Println(ebt)
	// fmt.Println(profit)
	// fmt.Println(ratio)

	// sum := 0
	// for   {
	// 	sum++
	// 	fmt.Println("Final Sum:", sum)
	// }


	// student1:= Student{"Alice", 20, "A"}

	// showDetails(&student1);	
	// fmt.Println("Student Name in main:", student1.Name)

	// job := struct{
	// 	Title string
	// 	Salary float64
	// }{
	// 	Title: "Software Engineer",
	// 	Salary: 75000.00,
	// }

	// fmt.Println("Job Title:", job.Title)
	// fmt.Println("Job Salary:", job.Salary)


	//Maps Demo
	studentTable := map[string]int{}

	studentTable["Alice"] = 20
	studentTable["Bob"] = 22
	studentTable["Charlie"] = 19

	fmt.Println("Student Table:", studentTable)
	
	number, exists := studentTable["bob"]
	fmt.Println("Bob's Age:", number, "Exists:", exists)

	for name, value := range studentTable {
		fmt.Println("Name:", name, "Value:", value)
	}


	// numbers := []int{1,2,3,4,5,6,7,8,9,10}

	// for i, value := range numbers {
	// 	fmt.Println("Index:", i, "Value:", value)
	// }

}

func showDetails(s1 *Student) { 
	s1.Name = "Mayank"
	fmt.Println("Student Name in the function:", s1.Name)
	fmt.Println("Student Age:", s1.Age)
	fmt.Println("Student Grade:", s1.Grade)	
}

func getUserInput(userInfo string) float64{
	var input float64
	fmt.Print("Enter ", userInfo, ": ")
	fmt.Scan(&input)

	return input
}