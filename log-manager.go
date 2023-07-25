package main

import "context"

type apiLogManager struct {
	store logStore
}

func newAPILogManager(ls logStore) *apiLogManager {
	return &apiLogManager{ls}
}

func (l *apiLogManager) SaveLog(ctx context.Context, r *logRecord) error {
	return l.store.SaveLog(ctx, r)
}

func (l *apiLogManager) GetLogByID(ctx context.Context, id string) (*logRecord, error) {
	return l.store.GetLogByID(ctx, id)
}

func (l *apiLogManager) FindLogs(ctx context.Context, filter logFilter) ([]*logRecord, error) {
	return l.store.FindLogs(ctx, filter)
}
