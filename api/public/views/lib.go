package views

import "time"

type Timestamp uint

func TimestampView(t time.Time) Timestamp {
	return Timestamp(t.Unix())
}

func NullableTimestampView(t *time.Time) *Timestamp {
	if t == nil {
		return nil
	}
	timestamp := Timestamp(t.Unix())
	return &timestamp
}
