package null

import (
	"encoding/json"
	"strconv"
)

// Interface guards for nullable integers.
var (
	_ json.Marshaler   = (Int)(nil)
	_ json.Unmarshaler = (Int)(nil)
)

// Int is a nullable integer value.
type Int = *IntData

// IntData contains data used to support nullable ints.
type IntData struct {
	Val  int
	Init bool
}

// ZeroInt is a null int.
var ZeroInt = &IntData{}

// NewInt creates a new nullable int.
func NewInt(i int) Int {
	return &IntData{
		Val:  i,
		Init: true,
	}
}

// MarshalJSON marshals an integer value sent by Discord. It implements the
// [json.Marshaler] interface.
func (i IntData) MarshalJSON() ([]byte, error) {
	if !i.Init {
		return []byte("null"), nil
	}

	return []byte(strconv.FormatInt(int64(i.Val), 10)), nil
}

// UnmarshalJSON unmarshals an integer value sent by Discord. It implements the
// [json.Unmarshaler] interface.
func (i *IntData) UnmarshalJSON(b []byte) error {
	s := string(b)
	if s == "null" {
		*i = *ZeroInt
		return nil
	}

	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}

	i.Val = int(v)
	i.Init = true
	return nil
}
