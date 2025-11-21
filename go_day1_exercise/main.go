package main

import(
	"fmt"
	"runtime"
	"os/user"
	""
	
) 

func main() {
	fmt.Println("Hello Motadata")
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.Version())
	fmt.Println("Number of CPUs:", runtime.NumCPU())
	fmt.Println("Number of Goroutines:", runtime.NumGoroutine())

	// Example YAML marshal/unmarshal usage
	type Person struct {
		Name string `yaml:"name"`
		Age  int    `yaml:"age"`
	}

	p := Person{Name: "Alice", Age: 30}
	data, err := yaml.Marshal(&p)
	if err != nil {
		fmt.Println("YAML Marshal error:", err)
	} else {
		fmt.Println("YAML output:", string(data))
	}

	var p2 Person
	err = yaml.Unmarshal(data, &p2)
	if err != nil {
		fmt.Println("YAML Unmarshal error:", err)
	} else {
		fmt.Printf("Unmarshaled struct: %+v\n", p2)
	}

	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("Error fetching user:", err)
	} else {
		fmt.Println("User:", currentUser.Username)
	}
	
}