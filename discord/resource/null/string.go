package null

import "encoding/json"

// Interface guards for nullable strings.
var (
	_ json.Marshaler   = (String)(nil)
	_ json.Unmarshaler = (String)(nil)
)

// String is a nullable string value.
type String = *StringData

// StringData contains data used to support nullable strings.
type StringData struct {
	Val  string
	Init bool
}

// ZeroString is a null string.
var ZeroString = &StringData{}

// NewString creates a new nullable string.
func NewString(s string) String {
	return &StringData{
		Val:  s,
		Init: true,
	}
}

// MarshalJSON marshals a string sent to Discord. It implements the
// [json.Marshaler] interface.
func (s StringData) MarshalJSON() ([]byte, error) {
	if !s.Init {
		return []byte("null"), nil
	}

	return json.Marshal(s.Val)
}

// UnmarshalJSON unmarshals a string sent by Discord. It implements the
// [json.Unmarshaler] interface.
func (s *StringData) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		*s = *ZeroString
		return nil
	}

	s.Init = true
	return json.Unmarshal(b, &s.Val)
}
