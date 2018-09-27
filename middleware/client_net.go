package middleware

import (
	"context"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

// CheckClientIsLocal checks that request comes from the local networks
func CheckClientIsLocal(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	localNetworks := []*net.IPNet{
		&net.IPNet{IP: net.IP{0x7f, 0x0, 0x0, 0x0}, Mask: net.IPMask{0xff, 0x0, 0x0, 0x0}},
		&net.IPNet{IP: net.IP{0xa, 0x0, 0x0, 0x0}, Mask: net.IPMask{0xff, 0x0, 0x0, 0x0}},
		&net.IPNet{IP: net.IP{0x64, 0x40, 0x0, 0x0}, Mask: net.IPMask{0xff, 0xc0, 0x0, 0x0}},
		&net.IPNet{IP: net.IP{0xac, 0x10, 0x0, 0x0}, Mask: net.IPMask{0xff, 0xf0, 0x0, 0x0}},
		&net.IPNet{IP: net.IP{0xc0, 0xa8, 0x0, 0x0}, Mask: net.IPMask{0xff, 0xff, 0x0, 0x0}},
	}
	p, ok := peer.FromContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.PermissionDenied, "peer info not-found")
	}
	addr, ok := p.Addr.(*net.TCPAddr)
	if !ok {
		return nil, status.Errorf(codes.PermissionDenied, "broken peer info")
	}
	for _, n := range localNetworks {
		if n.Contains(addr.IP.To4()) {
			return handler(ctx, req)
		}
	}
	return nil, status.Errorf(codes.PermissionDenied, "Request must be from local. Your IP: %s", addr.IP)
}
