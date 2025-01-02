package rate

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

const (
	key          = "X-RateLimit-"
	keyLimit     = key + "Limit"
	keyRemaining = key + "Remaining"
	keyReset     = key + "Reset"
	keyBucket    = key + "Bucket"
)

var ErrMissingValue = errors.New("rate: header missing value")

type Bucket struct {
	id        string
	limit     int
	remaining int
	reset     time.Time
	global    bool
}

// NewBucket returns an empty bucket.
func NewBucket() *Bucket {
	return new(Bucket)
}

func NewBucketFromHeader(h http.Header) (*Bucket, error) {
	limitString := h.Get(keyLimit)
	remainingString := h.Get(keyRemaining)
	resetString := h.Get(keyReset)
	bucket := h.Get(keyBucket)

	if limitString == "" {
		return nil, fmt.Errorf("%w: %s", ErrMissingValue, keyLimit)
	}

	if remainingString == "" {
		return nil, fmt.Errorf("%w: %s", ErrMissingValue, keyRemaining)
	}

	if resetString == "" {
		return nil, fmt.Errorf("%w: %s", ErrMissingValue, keyReset)
	}

	if bucket == "" {
		return nil, fmt.Errorf("%w: %s", ErrMissingValue, keyBucket)
	}

	limit, err := strconv.ParseInt(limitString, 10, 64)
	if err != nil {
		return nil, err
	}

	remaining, err := strconv.ParseInt(remainingString, 10, 64)
	if err != nil {
		return nil, err
	}

	resetValue, err := strconv.ParseInt(resetString, 10, 64)
	if err != nil {
		return nil, err
	}

	reset := time.Unix(resetValue, 0)

	b := &Bucket{
		limit:     int(limit),
		remaining: int(remaining),
		reset:     reset,
		id:        bucket,
	}

}
