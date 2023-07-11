package main

type apiLogManager struct {
	store    logStore
	latestID int
}

func newAPILogManager(ls logStore) *apiLogManager {
	return &apiLogManager{ls, 0}
}
