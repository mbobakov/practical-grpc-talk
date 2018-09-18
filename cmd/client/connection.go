package main

import (
	dynbalance "github.com/mbobakov/grpc-dynamic-balancer"
	"github.com/mbobakov/grpc-dynamic-balancer/provider/consul"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer"
)

func connect(tpe, target, sname string) (*grpc.ClientConn, error) {
	switch tpe {
	case "consul":
		bb := dynbalance.NewBalancerBuilder(sname, &consul.Consul{})
		balancer.Register(bb)
		return grpc.Dial(target,
			grpc.WithInsecure(),
			grpc.WithBalancerName(sname),
		)
	case "direct":
		return grpc.Dial(target,
			grpc.WithBlock(),
			grpc.WithInsecure(),
			grpc.WithBalancer(
				grpc.RoundRobin(
					NewPseudoResolver(target),
				),
			),
		)
	default:
		return nil, errors.Errorf("Unknow connection type: %s", tpe)
	}

}
