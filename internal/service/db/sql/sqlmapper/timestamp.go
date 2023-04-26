package sqlmapper

import "time"

func MapTimestampToTime(timestamp *string) (*time.Time, error) {
	if timestamp == nil {
		return nil, nil
	}

	t, err := time.Parse(time.RFC3339, *timestamp)
	if err != nil {
		return nil, err
	}
	return &t, nil
}
