package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	printUser := false

	flag.BoolVar(&printUser, "print-user", true, "Print a name if set to true")

	flag.Parse()

	user := os.Getenv("USER")

	if printUser {
		fmt.Println("Hello " + user)

		fmt.Printf("Hello second string %s\n", user)
	}
}
