package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/mbobakov/practical-grpc-talk/internal/server"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/grpclog"
)

func main() {
	var opts = struct {
		GRPCListen  string `long:"grpc.listen" env:"GRPC_LISTEN" default:":50051" description:"GRPC server interface"`
		DebugListen string `long:"debug.listen" env:"DEBUG_LISTEN" default:":6060" description:"Interface for serve debug information(metrics/health/pprof)"`
		Verbose     bool   `short:"v" env:"VERBOSE" description:"Enable Verbose log  output"`
	}{}

	_, err := flags.Parse(&opts)
	if err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		}
		os.Exit(1)
	}

	if opts.Verbose {
		logrus.SetLevel(logrus.DebugLevel)
	}

	grpclog.SetLoggerV2(&grpcLog{logrus.StandardLogger()})

	gr, ctx := errgroup.WithContext(context.Background())

	gr.Go(func() error {
		return server.ServeGRPC(ctx, opts.GRPCListen)
	})
	gr.Go(func() error {
		errCh := make(chan error)
		go func() {
			http.Handle("/metrics", promhttp.Handler())
			errCh <- http.ListenAndServe(opts.DebugListen, nil)
		}()
		select {
		case <-ctx.Done():
			return nil
		case err := <-errCh:
			return err
		}
	})

	if err := gr.Wait(); err != nil {
		log.Fatal(err)
	}
}

type grpcLog struct {
	*logrus.Logger
}

func (l *grpcLog) V(lvl int) bool {
	return true
}
