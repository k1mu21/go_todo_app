package main

import (
	"context"
	"log"
	"net" // Add this import statement
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Printf("need port number\n")
		os.Exit(1)
	}
	p := os.Args[1]
	l, err := net.Listen("tcp", ":"+p) // Fix the method name to "Listen"
	if err != nil {
		log.Fatalf("failed to listen port %s: %v", p, err) // Fix the typo "faild" to "failed"
	}
	if err := run(context.Background(), l); err != nil {
		log.Printf("failed to terminate: %v", err) // Fix the typo "faild" to "failed"
		os.Exit(1)
	}
}
