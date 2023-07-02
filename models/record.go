package models

import (
	"time"
)

type Record struct {
	ID       int
	Start    time.Time
	Stop     time.Time
	Duration int
	Project  string
}
