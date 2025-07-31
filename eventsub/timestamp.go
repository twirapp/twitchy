package eventsub

import (
	"fmt"
	"strings"
	"time"
)

// Timestamp is a time.Time wrapper to handle timestamps in RFC3339 format in nanoseconds (as all eventsub timestamps
// in this format).
type Timestamp struct {
	time.Time
}

// timestampFromString ...
func timestampFromString(timestamp string) (Timestamp, error) {
	timestamp = strings.Trim(timestamp, "\"")

	rfc3339NanoTimestamp, err := time.Parse(time.RFC3339Nano, timestamp)
	if err != nil {
		return Timestamp{}, fmt.Errorf("parse to rfc3339 in nanoseconds: %w", err)
	}

	return Timestamp{
		Time: rfc3339NanoTimestamp,
	}, nil
}

func (t *Timestamp) UnmarshalJSON(payload []byte) error {
	rfc3339NanoTimestamp, err := timestampFromString(string(payload))
	if err != nil {
		return fmt.Errorf("timestamp from string: %w", err)
	}

	t.Time = rfc3339NanoTimestamp.Time
	return nil
}

// TimestampUTC is a time.Time wrapper to handle timestamps in RFC3339 format in nanoseconds with UTC location set.
type TimestampUTC struct {
	time.Time
}

// timestampUTCFromString parses specified string-based timestamp to the TimestampUTC.
func timestampUTCFromString(timestamp string) (TimestampUTC, error) {
	rfc3339NanoTimestamp, err := timestampFromString(timestamp)
	if err != nil {
		return TimestampUTC{}, fmt.Errorf("timestamp from string: %w", err)
	}

	return TimestampUTC{
		Time: rfc3339NanoTimestamp.Time.UTC(),
	}, nil
}

func (tu *TimestampUTC) UnmarshalJSON(payload []byte) error {
	timestamp, err := timestampUTCFromString(string(payload))
	if err != nil {
		return fmt.Errorf("timestamp utc from string: %w", err)
	}

	tu.Time = timestamp.Time
	return nil
}
