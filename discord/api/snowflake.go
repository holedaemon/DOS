package api

import "time"

// Epoch is the first second of 2015.
const Epoch = 1420070400000

// Snowflake represents a unique identifier utilized by Discord's API.
// Refer to the [developer documentation] for more information.
//
// [developer documentation]: https://discord.com/developers/docs/reference#snowflakes
type Snowflake uint64

// Timestamp returns a [time.Time] of when the [Snowflake] was created.
func (s Snowflake) Timestamp() time.Time {
	ts := int64((s >> 22) + Epoch)
	return time.UnixMilli(ts)
}

// WorkerID returns the internal worker ID of the [Snowflake].
func (s Snowflake) WorkerID() uint8 {
	return uint8((s & 0x3E0000) >> 17)
}

// ProcessID returns the internal process ID of the [Snowflake].
func (s Snowflake) ProcessID() uint8 {
	return uint8((s & 0x1F000) >> 12)
}

// Increment returns the [Snowflake] increment.
func (s Snowflake) Increment() uint16 {
	return uint16(s & 0xFFF)
}
