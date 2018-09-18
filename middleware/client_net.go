package middleware

import (
	"context"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

// CheckClientIsLocal checks that request comes from 127.0.0.1
func CheckClientIsLocal(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	if err := checkNetFromContext(ctx, "127.0.0.1"); err != nil {
		return nil, err
	}
	return handler(ctx, req)
}

// CheckClientIsLocalStream checks peer IP-address. It should be 127.0.0.1
func CheckClientIsLocalStream(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	if err := checkNetFromContext(ss.Context(), "127.0.0.1"); err != nil {
		return err
	}
	grpclog.Infof("%#v", srv)
	grpclog.Infof("%#v", info)
	grpclog.Infof("Started")
	err := handler(srv, ss)
	grpclog.Infof("Finished %v", err)
	return nil
}

func checkNetFromContext(ctx context.Context, nt string) error {
	p, ok := peer.FromContext(ctx)
	if !ok {
		return status.Errorf(codes.PermissionDenied, "peer info not-found")
	}
	addr, ok := p.Addr.(*net.TCPAddr)
	if !ok {
		return status.Errorf(codes.PermissionDenied, "broken peer info")
	}
	if addr.IP.String() != nt {
		return status.Errorf(codes.PermissionDenied, "Request must be from %s", nt)
	}
	return nil
}
