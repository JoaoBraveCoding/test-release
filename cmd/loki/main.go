package main

import (
	"fmt"

	"github.com/hashicorp/go-uuid"
)

func main() {
	fmt.Println("New feature for 0.1.0")
	fmt.Println("New feature for 0.2.0")
	fmt.Println("New feature for 0.2.0")
	fmt.Println("New feature for 0.2.0")
	fmt.Println("New feature for 3.0.0")

	uid, err := uuid.GenerateUUID()
	if err != nil {
		fmt.Println("Error generating UUI")
	}
	fmt.Println("This is not really Loki, just meant to look like it. Here's a UUID: ", uid)
}
