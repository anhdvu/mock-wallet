package main

type apiLogManager struct {
	store logStore
}

func newAPILogManager(ls logStore) *apiLogManager {
	return &apiLogManager{ls}
}
