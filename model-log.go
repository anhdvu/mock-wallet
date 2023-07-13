package main

import (
	"time"

	"github.com/gofrs/uuid/v5"
)

type record struct {
	Checksum     string
	Type         string
	Time         time.Time
	RawRequest   []byte
	JSON         []byte
	KLVBreakdown [][4]string
	Response     []byte
	ID           uuid.UUID
}
