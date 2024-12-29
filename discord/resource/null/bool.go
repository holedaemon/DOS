package null

import (
	"encoding/json"
	"strconv"
)

// Interface guards for null booleans.
var (
	_ json.Marshaler   = (Bool)(nil)
	_ json.Unmarshaler = (Bool)(nil)
)

// Bool is a nullable boolean value.
type Bool = *BoolData

// BoolData contains data used to support nullable booleans.
type BoolData struct {
	Val  bool
	Init bool
}

var (
	// ZeroBool is a null bool.
	ZeroBool = &BoolData{}

	// True is a valid bool that is true.
	True = &BoolData{
		Val:  true,
		Init: true,
	}

	// False is a valid bool that is false.
	False = &BoolData{
		Val:  false,
		Init: true,
	}
)

// MarshalJSON marshals a boolean value sent to Discord. It implements the
// [json.Marshaler] interface.
func (b BoolData) MarshalJSON() ([]byte, error) {
	if !b.Init {
		return []byte("null"), nil
	}

	return []byte(strconv.FormatBool(b.Val)), nil
}

// UnmarshalJSON unmarshals a boolean value sent by Discord. It implements the
// [json.Unmarshaler] interface.
func (b *BoolData) UnmarshalJSON(v []byte) error {
	s := string(v)

	if s == "null" {
		*b = *ZeroBool
		return nil
	}

	val, err := strconv.ParseBool(s)
	if err != nil {
		return err
	}

	b.Val = val
	b.Init = true
	return nil
}
