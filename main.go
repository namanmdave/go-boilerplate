package main

import (
	"go-boilerplate/cmd"
	"log"
)

func main() {
	if err := cmd.Init(); err != nil {
		log.Fatalf("Application failed: %v", err)
	}
}
