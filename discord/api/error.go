package api

import (
	"strconv"
	"strings"
)

// ErrorCode is a JSON error code sent by Discord's API. A complete list of codes can
// be found on the [developer documentation].
//
// [developer documentation]: https://discord.com/developers/docs/topics/opcodes-and-status-codes#json
type ErrorCode int

// Error represents an error returned by Discord's API. It implements the error
// interface. Refer to the [developer documentation] for more information.
//
// [developer documentation]: https://discord.com/developers/docs/reference#error-messages
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`

	// TODO: Handle the errors field.
}

// Error returns a human readable error.
func (e *Error) Error() string {
	var sb strings.Builder

	sb.WriteString("discord: api: ")

	if e.Code > 0 {
		sb.WriteString(strconv.Itoa(e.Code))
	}

	if e.Message != "" {
		if e.Code > 0 {
			sb.WriteString(": ")
		}

		sb.WriteString(e.Message)
	}

	return sb.String()
}
