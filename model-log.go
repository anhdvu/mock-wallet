package main

import (
	"time"

	"github.com/gofrs/uuid/v5"
)

type logRecord struct {
	Checksum     string
	Type         string
	Time         time.Time
	RawRequest   string
	JSON         string
	KLVBreakdown [][4]string
	Response     string
	ID           uuid.UUID
}

type logFilter struct {
	offset int
	limit  int
}

func newLogRecord(t string) (*logRecord, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	return &logRecord{
		Type: t,
		Time: time.Now(),
		ID:   id,
	}, nil
}
