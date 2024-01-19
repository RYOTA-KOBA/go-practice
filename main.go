package main

import (
	"fmt"
)

func main() {
	firstName, lastName := "John", "Doe"
	age := 32
	const HTTPStatusOK = 200
	fmt.Println(firstName, lastName, age, "status", HTTPStatusOK)
}
