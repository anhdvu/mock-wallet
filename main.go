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
	port := getEnvOrDefault("PORT", devPort)

	cfg := config{
		env:  "dev",
		port: fmt.Sprintf(":%s", port),
	}

	app := newApp(cfg)

	rc, err := openRedis("rediss://default:b53dd65dcb914e96969ab06c5ca625d1@upward-tahr-33126.upstash.io:33126")
	if err != nil {
		app.logger.Fatal(err)
		os.Exit(1)
	}

	apiLogger := newAPILogManager(newRedisLogStore(rc))
	app.apiLogger = apiLogger

	err = app.run()
	if err != nil {
		trace := debug.Stack()
		log.Fatalf("%s\n%s", err, trace)
	}
}
