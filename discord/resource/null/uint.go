package null

import (
	"encoding/json"
	"strconv"
)

// Interface guards for nullable uints.
var (
	_ json.Marshaler   = (Uint)(nil)
	_ json.Unmarshaler = (Uint)(nil)
)

// Uint is a nullable uint value.
type Uint = *UintData

// UintData contains data used to support nullable uints.
type UintData struct {
	Val  uint
	Init bool
}

// ZeroUint is a null uint.
var ZeroUint = &UintData{}

// NewUint creates a new nullable uint.
func NewUint(i uint) Uint {
	return &UintData{
		Val:  i,
		Init: true,
	}
}

// MarshalJSON marshals a uint value sent to Discord. It implements the
// [json.Marshaler] interface.
func (u UintData) MarshalJSON() ([]byte, error) {
	if !u.Init {
		return []byte("null"), nil
	}

	return []byte(strconv.FormatUint(uint64(u.Val), 10)), nil
}

// UnmarshalJSON unmarshals a uint value sent by Discord. It implements the
// [json.Unmarshaler] interface.
func (u *UintData) UnmarshalJSON(b []byte) error {
	s := string(b)

	if s == "null" {
		*u = *ZeroUint
		return nil
	}

	v, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return err
	}

	u.Val = uint(v)
	u.Init = true
	return nil
}
