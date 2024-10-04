package main

import (
	"MakeAPackage/makeapackage"
	"log"
)

func main() {
	log.SetPrefix("main: ")
	log.SetFlags(0)

	result, err := makeapackage.MakeAPackage("Gae")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	log.Printf("Result: %s", result)
}
