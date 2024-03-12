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

	uid, err := uuid.GenerateUUID()
	if err != nil {
		fmt.Println("Error generating UUID")
	}
	fmt.Println("This is not really the Loki Operator, just meant to look like it. Here's a UUID: ", uid)
}
