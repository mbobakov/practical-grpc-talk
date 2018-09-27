package server

import (
	"context"
	"net"
	"time"

	"github.com/golang/protobuf/ptypes/duration"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	apiV1 "github.com/mbobakov/practical-grpc-talk/api/v1"
	"github.com/mbobakov/practical-grpc-talk/middleware"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	hlzpb "google.golang.org/grpc/health/grpc_health_v1"
)

// TM is a time machine server engine
type TM struct{}

// ServeGRPC on the provided address.
// Blocks until stop or error not occured
func ServeGRPC(ctx context.Context, l string) error {
	lis, err := net.Listen("tcp", l)
	if err != nil {
		return errors.Wrap(err, "Could't start listen")
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
				grpc_logrus.UnaryServerInterceptor(logrus.NewEntry(logrus.StandardLogger())),
				grpc_prometheus.UnaryServerInterceptor,
				grpc_validator.UnaryServerInterceptor(),
				middleware.CheckClientIsLocal,
			),
		),
	)

	hlzpb.RegisterHealthServer(s, health.NewServer())
	apiV1.RegisterTimeMachineServer(s, &TM{})

	e := make(chan error)
	go func() {
		e <- s.Serve(lis)
	}()
	select {
	case <-ctx.Done():
		s.GracefulStop()
		return nil
	case err := <-e:
		return err
	}
}

// Jump to the concrete moment in time
func (tm *TM) Jump(ctx context.Context, jr *apiV1.JumpRequest) (*duration.Duration, error) {
	n := time.Now()
	sd := n.Unix() - jr.To.GetSeconds()
	nd := n.UnixNano() - n.Unix()*int64(time.Second) - int64(jr.To.GetNanos())
	return &duration.Duration{
		Seconds: sd,
		Nanos:   int32(nd),
	}, nil
}
