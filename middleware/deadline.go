package middleware

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CheckDeadlineIsSet checks that request comes with propper deadline
func CheckDeadlineIsSet(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	d, ok := ctx.Deadline()
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "Deadline must be set for a request")
	}
	timeout := d.UTC().Sub(time.Now().UTC())
	if timeout < 5*time.Second || timeout > 60*time.Second {
		return nil, status.Errorf(codes.InvalidArgument, "Deadline must be in 5-60 sec range")

	}
	return handler(ctx, req)
}
