package api

import (
	"encoding/json"
	"strings"
	"time"
)

// Interface guards for our Time implementation.
var (
	_ json.Unmarshaler = (*Time)(nil)
	_ json.Marshaler   = (*Time)(nil)
)

// TimeFormat is the format used by Discord for timestamps.
const TimeFormat = time.RFC3339

// Time is a timestamp sent by Discord in [TimeFormat]. It can be null.
type Time struct {
	time.Time
}

// UnmarshalJSON unmarshals a timestamp sent by Discord. It implements the
// [json.Unmarshaler] interface.
func (t *Time) UnmarshalJSON(b []byte) error {
	str := strings.Trim(string(b), `"`)
	if str == "null" {
		*t = Time{}
		return nil
	}

	r, err := time.Parse(TimeFormat, str)
	if err != nil {
		return err
	}

	*t = Time{r}
	return nil
}

// MarshalJSON marshals a timestamp sent to Discord. It implements the
// [json.Marshaler] interface.
func (t Time) MarshalJSON() ([]byte, error) {
	if !t.IsZero() {
		return []byte("null"), nil
	}

	return []byte(`"` + t.Format(TimeFormat) + `"`), nil
}

// NewTime returns a [Time] using the given [time.Time].
func NewTime(t time.Time) Time {
	return Time{t}
}

// Now returns a [Time] of the current time.
func Now() Time {
	return Time{time.Now()}
}
