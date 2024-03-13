package main

import (
	"fmt"

	"github.com/hashicorp/go-uuid"
)

func main() {
	
	fmt.Println("New feature for 0.1.0")
	fmt.Println("New feature for 0.1.1")
	fmt.Println("New feature for 0.3.0")
	fmt.Println("New feature for 0.3.1")
	fmt.Println("New feature for 0.3.2")
	fmt.Println("New feature for 0.4.0")
	fmt.Println("New feature for 0.4.1")
	fmt.Println("New feature for 0.5.0")
	fmt.Println("New feature for 0.6.0")
	fmt.Println("New feature for 0.7.0")
	fmt.Println("New feature for 0.8.0")
	fmt.Println("New feature for 0.9.1")
	fmt.Println("New feature for 0.10.0")
	fmt.Println("New feature for 0.11.0")
	fmt.Println("New feature for 0.12.0")
	fmt.Println("New feature for 0.13.0")
	fmt.Println("New feature for 0.12.0")

	uid, err := uuid.GenerateUUID()
	if err != nil {
		fmt.Println("Error generating UUID")
	}
	fmt.Println("This is not really the Loki Operator, just meant to look like it. Here's a UUID: ", uid)
}
