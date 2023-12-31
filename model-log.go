package main

import (
	"time"

	"github.com/gofrs/uuid/v5"
)

type logRecord struct {
	Time         time.Time
	Checksum     string
	Type         string
	RawRequest   string
	JSON         string
	Response     string
	KLVBreakdown [][4]string
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

func (l *logRecord) addInfo(checksumString, jsonString, klvString string) error {
	l.Checksum = checksumString

	l.JSON = jsonString

	klv, err := breakDownKLV(klvString)
	if err != nil {
		l.KLVBreakdown = nil
	} else {
		l.KLVBreakdown = klv
	}

	return err
}
