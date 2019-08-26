package model

import "time"

type NullString struct {
	String string
	Valid  bool
}

type NullTime struct {
	Time  time.Time
	Valid bool
}
