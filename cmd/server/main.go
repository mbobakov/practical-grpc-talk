package main

import (
	"context"
	"log"
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/mbobakov/practical-grpc-talk/internal/server"
)

func main() {
	var opts = struct {
		GRPCListen  string `long:"grpc.listen" env:"GRPC_LISTEN" default:":50501" description:"GRPC server interface"`
		DebugListen string `long:"debug.listen" env:"DEBUG_LISTEN" default:":6060" description:"Interface for serve debug information(metrics/health/pprof)"`
		Verbose     bool   `long:"v" env:"VERBOSE" description:"Enable Verbose log  output"`
	}{}

	_, err := flags.Parse(&opts)
	if err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		}
		os.Exit(1)
	}
	err = server.ServeGRPC(context.Background(), opts.GRPCListen)
	if err != nil {
		log.Fatal(err)
	}

}
