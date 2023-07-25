package main

import "os"

func getEnvOrDefault(env, def string) string {
	v, ok := os.LookupEnv(env)
	if !ok {
		v = def
	}

	return v
}
