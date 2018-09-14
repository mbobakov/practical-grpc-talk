package server

import (
	"context"
	"net"
	"time"

	"github.com/golang/protobuf/ptypes/duration"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/mbobakov/practical-grpc-talk/api"
	"github.com/mbobakov/practical-grpc-talk/middleware"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// TS is a time server
// Implements api.TimeServer
type TS struct{}

// ServeGRPC interface on the provided address.
// Blocks until stop or error not occured
func ServeGRPC(ctx context.Context, l string) error {
	lis, err := net.Listen("tcp", l)
	if err != nil {
		return errors.Wrap(err, "Could't start listen")
	}
	s := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.CheckClientIsLocal),
		grpc.StreamInterceptor(middleware.CheckClientIsLocalStream),
	)
	api.RegisterTimeServer(s, &TS{})
	errCh := make(chan error, 1)
	go func() {
		errCh <- s.Serve(lis)
	}()
	select {
	case err := <-errCh:
		return errors.Wrap(err, "GRPC server error")
	case <-ctx.Done():
		s.GracefulStop()
		return nil
	}
}

// CurrentDayLength returns duration of current day
// ! works in  the server TimeZone
func (t *TS) CurrentDayLength(ctx context.Context, e *empty.Empty) (*duration.Duration, error) {
	now := time.Now()
	year, month, day := now.Date()
	durFrom := time.Since(time.Date(year, month, day, 0, 0, 0, 0, now.Location()))
	return &duration.Duration{
			Seconds: int64(durFrom.Seconds()),
			Nanos:   int32(durFrom.Nanoseconds() - int64(durFrom.Seconds())*int64(time.Second)),
		},
		nil
}

// Clock tick every second
func (t *TS) Clock(req *empty.Empty, st api.Time_ClockServer) error {
	tickr := time.NewTicker(time.Second)
	for tm := range tickr.C {
		err := st.Send(&timestamp.Timestamp{
			Seconds: int64(tm.Unix()),
			Nanos:   int32(tm.UnixNano() - int64(tm.Unix())*int64(time.Second)),
		})
		if err != nil {
			return err
		}
	}
	return nil
}
