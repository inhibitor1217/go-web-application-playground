package views

import "time"

type Timestamp uint

func TimestampView(t time.Time) Timestamp {
	return Timestamp(t.Unix())
}
