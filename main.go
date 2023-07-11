package main

import (
	"log"
	"runtime/debug"
)

func main() {
	cfg := config{
		env:  "dev",
		port: 8000,
	}

	app := New(cfg)

	err := app.run()
	if err != nil {
		trace := debug.Stack()
		log.Fatalf("%s\n%s", err, trace)
	}
}
