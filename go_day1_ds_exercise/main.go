package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var seats [300]int

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n--- Seat Booking System ---")
		fmt.Println("1. Show available seats")
		fmt.Println("2. Book a seat")
		fmt.Println("3. Exit")
		fmt.Print("Enter choice: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			showAvailableSeats(seats)

		case "2":
			fmt.Print("Enter seat number to book (1–300): ")
			scanner.Scan()
			input := strings.TrimSpace(scanner.Text())

			seatNum, err := strconv.Atoi(input)
			if err != nil || seatNum < 1 || seatNum > 300 {
				fmt.Println("Invalid seat number.")
				continue
			}

			if seats[seatNum-1] == 1 {
				fmt.Println("Seat already booked!")
			} else {
				seats[seatNum-1] = 1
				fmt.Printf("Seat %d booked successfully.\n", seatNum)
			}

		case "3":
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func showAvailableSeats(seats [300]int) {
	fmt.Println("\nAvailable seats:")
	for i := 0; i < len(seats); i++ {
		if seats[i] == 0 {
			fmt.Printf("%d ", i+1)
		}
	}
	fmt.Println()
}
