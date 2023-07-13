package main

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"
)

const (
	devPort = "8080"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = devPort
	}

	cfg := config{
		env:  "dev",
		port: fmt.Sprintf(":%s", port),
	}

	app := New(cfg)

	err := app.run()
	if err != nil {
		trace := debug.Stack()
		log.Fatalf("%s\n%s", err, trace)
	}
}
