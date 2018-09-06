package server

import (
	"context"

	"github.com/golang/protobuf/ptypes/duration"
	"github.com/golang/protobuf/ptypes/empty"
)

// TS is a time server
// Implements api.TimeServer
type TS struct{}

// ServeGRPC interface on the provided address.
// Blocks until stop or error not occured
func ServeGRPC(ctx context.Context, l string) error {
	return nil
}

// CurrentDayLength returns duration of current day
// ! works in  the server TimeZone
func (t *TS) CurrentDayLength(ctx context.Context, e *empty.Empty) (*duration.Duration, error) {
	panic("NotImpemented!")
}
