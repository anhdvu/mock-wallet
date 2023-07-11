package main

import "context"

type logStore interface {
	Write(ctx context.Context, r record) error
	All(ctx context.Context) ([]record, error)
	Recent(ctx context.Context) ([]record, error)
}
