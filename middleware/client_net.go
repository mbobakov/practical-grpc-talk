package middleware

import (
	"context"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

// CheckClientIsLocal checks that request comes from 127.0.0.1
func CheckClientIsLocal(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	if err := checkNetFromContext(ctx); err != nil {
		return nil, err
	}
	return handler(ctx, req)
}

// CheckClientIsLocalStream checks peer IP-address. It should be 127.0.0.1
func CheckClientIsLocalStream(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	if err := checkNetFromContext(ss.Context()); err != nil {
		return err
	}
	return handler(srv, ss)
}

func checkNetFromContext(ctx context.Context) error {

	localNetworks := []*net.IPNet{
		&net.IPNet{IP: net.IP{0xa, 0x0, 0x0, 0x0}, Mask: net.IPMask{0xff, 0x0, 0x0, 0x0}},
		&net.IPNet{IP: net.IP{0x64, 0x40, 0x0, 0x0}, Mask: net.IPMask{0xff, 0xc0, 0x0, 0x0}},
		&net.IPNet{IP: net.IP{0xac, 0x10, 0x0, 0x0}, Mask: net.IPMask{0xff, 0xf0, 0x0, 0x0}},
		&net.IPNet{IP: net.IP{0xc0, 0xa8, 0x0, 0x0}, Mask: net.IPMask{0xff, 0xff, 0x0, 0x0}},
	}

	p, ok := peer.FromContext(ctx)
	if !ok {
		return status.Errorf(codes.PermissionDenied, "peer info not-found")
	}
	addr, ok := p.Addr.(*net.TCPAddr)
	if !ok {
		return status.Errorf(codes.PermissionDenied, "broken peer info")
	}

	for _, n := range localNetworks {
		if n.Contains(addr.IP.To4()) {
			return nil
		}
	}

	return status.Errorf(codes.PermissionDenied, "Request must be from local. Your IP: %s", addr.IP)
}
