package main

import "context"

type logStore interface {
	GetLogByID(ctx context.Context, id string) (*logRecord, error)
	SaveLog(ctx context.Context, r *logRecord) error
	FindLogs(ctx context.Context, filter logFilter) ([]*logRecord, error)
}
