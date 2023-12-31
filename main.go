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
	redisConString := getEnvOrDefault("REDIS_CONNECTION", "redis://localhost:6379")

	cfg := config{
		env:  "dev",
		port: fmt.Sprintf(":%s", port),
	}

	app := newApp(cfg)

	rc, err := openRedis(redisConString)
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
