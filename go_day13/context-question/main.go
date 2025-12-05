package main

import (
	"context"
	"fmt"
	"time"
	// "fmt"
)

func main(){
	// var wg sync.WaitGroup
	fmt.Println("🚀 Application started.")
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()
	

	data, err := fetchData(ctx)

	if err != nil {
		fmt.Printf("🛑 Main function received error: %v\n", err)
	} else {
		fmt.Printf("✅ Main function received data: %s\n", data)
	}
	fmt.Println("🏁 Application finished.")

}

func fetchData(ctx context.Context) (string, error){
	fmt.Println("⏳ Starting data fetch operation (simulated 3s delay)...")

	result := make(chan string, 1)
	go func ()  {
		time.Sleep(4*time.Second)
		result <- "Hello from gor"
	}()


	for{
		select {
		case <-ctx.Done():
			fmt.Println("Context closed!!")
			return "", ctx.Err()
		case val := <-result :

			return val, nil
		}
	}
}