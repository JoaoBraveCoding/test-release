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
	fmt.Println("New feature for 0.7.1")
	fmt.Println("New feature for 0.8.0")
	fmt.Println("New feature for 0.9.0")
	fmt.Println("New feature for 0.9.1")
	fmt.Println("New feature for 0.9.2")
	fmt.Println("New feature for 0.9.3")
	fmt.Println("New feature for 0.10.0")
	fmt.Println("New feature for 0.10.1")
	fmt.Println("New feature for 0.10.2")
	fmt.Println("New feature for 0.10.3")
	fmt.Println("New feature for 0.10.4")
	fmt.Println("New feature for 0.10.2")
	fmt.Println("New feature for 0.10.3")
	fmt.Println("New feature for 0.10.4")
	fmt.Println("New feature for 0.10.5")
	fmt.Println("New feature for 0.10.5")
	fmt.Println("New feature for 0.10.6")
	fmt.Println("New feature for 0.10.7")
	fmt.Println("New feature for 0.10.8")
	fmt.Println("New feature for 0.10.9")
	fmt.Println("New feature for 0.10.10")
	fmt.Println("New feature for 0.10.11")
	fmt.Println("New feature for 0.11.0")
	fmt.Println("New feature for 0.11.1")
	fmt.Println("New feature for 0.11.2")
	fmt.Println("New feature for 0.11.3")
	fmt.Println("New feature for 0.11.4")

	
	uid, err := uuid.GenerateUUID()
	if err != nil {
		fmt.Println("Error generating UUID")
	}
	fmt.Println("This is not really the Loki Operator, just meant to look like it. Here's a UUID: ", uid)
}
