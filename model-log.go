package main

import "time"

type record struct {
	Type         string
	Time         time.Time
	RawRequest   []byte
	JSON         []byte
	KLVBreakdown [][4]string
	Response     []byte
	ID           int
}
